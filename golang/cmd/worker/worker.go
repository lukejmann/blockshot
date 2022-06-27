package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lukejmann/blockshot/golang/chain"
	"github.com/lukejmann/blockshot/golang/datagen"
	"github.com/lukejmann/blockshot/golang/env"
)

func main() {
	wrk, err := New()
	if err != nil {
		log.Fatalln("Unable to initialize worker", err)
	}

	log.Println("Starting worker...")
	wrk.Run()
}

type Worker struct {
	env        env.Env
	wg         sync.WaitGroup
	signalChan chan os.Signal
	exitChan   chan int
	datagen    *datagen.Datagen
	chain      *chain.Chain
}

func Connect(e env.Env) (dg *datagen.Datagen, c *chain.Chain, err error) {
	wssEndpoint := "wss://eth-mainnet.alchemyapi.io/v2/" + os.Getenv("ALCHEMY_API_KEY")
	fmt.Println("WSS endpoint:", wssEndpoint)
	c, err = chain.New(wssEndpoint)
	if err != nil {
		return nil, nil, fmt.Errorf("Error connecting to chain: %s", err)
	}

	zoraEndpoint := os.Getenv("ZORA_ENDPOINT")
	dg, err = datagen.New(e, zoraEndpoint)
	if err != nil {
		return nil, nil, fmt.Errorf("Error connecting to datagen: %s", err)
	}

	return dg, c, nil
}

func New() (*Worker, error) {
	godotenv.Load()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	env, err := env.New()
	if err != nil {
		return nil, err
	}

	dg, c, err := Connect(env)
	if err != nil {
		return nil, err
	}

	return &Worker{
		env,
		sync.WaitGroup{},
		signalChan,
		make(chan int),
		dg,
		c,
	}, nil
}

func (w *Worker) Run() {
	go func() {
		for {
			select {
			case s := <-w.signalChan:
				w.handleExit(s)
			}
		}
	}()

	bChan := make(chan *big.Int, 1)
	sub, err := w.chain.WatchNewBlock(bChan)
	if err != nil {
		log.Println("Unable to subscribe to new blocks", err)
		return
	}
	for {
		select {
		case e := <-sub.Err():
			panic(e)
		case newBlock := <-bChan:
			fmt.Printf("new block: %v\n", newBlock)
			w.doAsync(func() error {
				return w.datagen.FetchMintsAtBlock(int(newBlock.Int64()) - 10)
			})
			w.doAsync(func() error {
				return w.datagen.FixMissingMetadataFromBlock(int(newBlock.Int64()) - 12)
			})
		case code := <-w.exitChan:
			os.Exit(code)
		}
	}

}

func (w *Worker) handleExit(s os.Signal) {
	log.Println("Handling os signal...")

	w.wg.Wait()
	log.Println("Goroutines finished up")

	switch s {
	case syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT:
		log.Println("Received shutdown signal:", s)
		w.exitChan <- 0

	default:
		log.Println("Unknown signal!?", s)
		w.exitChan <- 1
	}
}

func (w *Worker) doAsync(fn func() error) {
	w.wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		err := fn()
		if err != nil {
			log.Println(err)
		}
	}(&w.wg)
}

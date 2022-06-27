package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

type Chain struct {
	client *ethclient.Client
}

func New(websocket string) (*Chain, error) {
	c, err := ethclient.Dial(websocket)
	if err != nil {
		return nil, err
	}
	return &Chain{
		client: c,
	}, nil
}

func (c *Chain) Client() *ethclient.Client {
	return c.client
}

// func (c *)
func (c *Chain) WatchNewBlock(sink chan<- *big.Int) (event.Subscription, error) {
	headS := make(chan *types.Header, 1)
	sub, err := c.client.SubscribeNewHead(context.Background(), headS)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case head := <-headS:
				sink <- head.Number
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/lukejmann/blockshot/golang/env"
	"github.com/lukejmann/blockshot/golang/errors"
	"github.com/lukejmann/blockshot/golang/server/write"
)

var isDev = false

func init() {
	isDev = os.Getenv("ENV") == "dev"
}

type server struct {
	env    env.Env
	router *httprouter.Router
}

func New() (*server, error) {
	env, err := env.New()
	if err != nil {
		return nil, err
	}

	srv := &server{
		env: env,
	}

	srv.ConfigureRouter()

	return srv, nil
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving request: ", r.URL.Path)
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head != "api" {
		write.Error(errors.RouteNotFound)
	}

	srv.router.ServeHTTP(w, r)
}

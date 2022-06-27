package server

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
	"github.com/lukejmann/blockshot/golang/env"
	"github.com/lukejmann/blockshot/golang/errors"
	"github.com/lukejmann/blockshot/golang/server/handlers"
	"github.com/lukejmann/blockshot/golang/server/write"
)

func (srv *server) ConfigureRouter() {
	srv.router = httprouter.New()

	srv.router.MethodNotAllowed = write.Error(errors.BadRequestMethod)
	srv.router.NotFound = write.Error(errors.RouteNotFound)
	srv.router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Println("Panic on", r.URL.Path)
		debug.PrintStack()
		write.Error(errors.InternalError)(w, r)
	}

	srv.GET("/block/:id", handlers.GetMintsAtBlock)
	srv.GET("/highest", handlers.GetHighestBlock)

}

type srvHandler func(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc

func (srv *server) GET(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodGet, path, srv.wrap(handler))
}
func (srv *server) PUT(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodPut, path, srv.wrap(handler))
}
func (srv *server) POST(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodPost, path, srv.wrap(handler))
}
func (srv *server) DELETE(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodDelete, path, srv.wrap(handler))
}

func (srv *server) wrap(h srvHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn := withEnv(srv.env, h, w, r)
		wrapped := lag(csrf(cors(fn)))
		wrapped(w, r)
	}
}

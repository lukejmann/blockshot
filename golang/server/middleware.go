package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/lukejmann/blockshot/golang/env"
	"github.com/lukejmann/blockshot/golang/errors"
	"github.com/lukejmann/blockshot/golang/server/write"
)

func withEnv(env env.Env, h srvHandler, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return h(env, w, r)
}

func lag(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			time.Sleep(time.Millisecond * 500)
		}
		fn(w, r)
	}
}

const localDev = "https://localhost"

func validateOrigin(r *http.Request) string {
	appRoot := os.Getenv("APP_ROOT")
	switch r.Header.Get("Origin") {
	case appRoot:
		return appRoot
	case localDev:
		return localDev
	default:
		return ""
	}
}

func csrf(fn http.HandlerFunc) http.HandlerFunc {
	fmt.Println("csrf")
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(r.URL.Path) {
			fn(w, r)
			return
		}

		if r.Method != http.MethodOptions {
			if r.Header.Get("Origin") != "" && validateOrigin(r) == "" {
				fn = write.Error(errors.BadOrigin)
			} else if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
				fn = write.Error(errors.BadCSRF)
			}
		}
		fn(w, r)
	}
}

func cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(r.URL.Path) {
			fn(w, r)
			return
		}

		if origin := validateOrigin(r); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

		if r.Method == http.MethodOptions {
			fn = write.Success()
		}
		fn(w, r)
	}
}

var bypassPaths = []string{}

func skipCorsAndCSRF(path string) bool {
	for _, c := range bypassPaths {
		if path == c {
			return true
		}
	}

	return false
}

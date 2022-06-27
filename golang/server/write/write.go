package write

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lukejmann/blockshot/golang/errors"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Error(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, code := errors.GetCode(err)
		if !found {
			log.Println("Unexpected Error: ", err)
			err = errors.InternalError
		}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&errorResponse{Error: err.Error()})
	}
}

func JSON(obj interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(obj)
	}
}

func JSONorErr(obj interface{}, err error) http.HandlerFunc {
	if err != nil {
		return Error(err)
	}

	return JSON(obj)
}

func SuccessOrErr(err error) http.HandlerFunc {
	if err != nil {
		return Error(err)
	}

	return Success()
}

func Success() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&map[string]bool{"success": true})
	}
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/lukejmann/blockshot/golang/env"
	"github.com/lukejmann/blockshot/golang/errors"
	"github.com/lukejmann/blockshot/golang/server/write"
)

func GetMintsAtBlock(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	block, err := getID(r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}

	mints, err := env.DB().GetMintsForBlock(r.Context(), int32(block))
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.MintsNotFound)
		}
		return write.Error(err)
	}

	return write.JSON(mints)
}

func GetHighestBlock(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	highest, err := env.DB().GetHighestBlock(r.Context())
	fmt.Println("highest", highest)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.InternalError)
		}
		return write.Error(err)
	}
	return write.JSON(highest)
}

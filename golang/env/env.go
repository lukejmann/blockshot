package env

import (
	"github.com/lukejmann/blockshot/golang/db/wrapper"
)

type Env interface {
	DB() wrapper.Querier
}

// default impl
type env struct {
	db wrapper.Querier
}

func (e *env) DB() wrapper.Querier {
	return e.db
}

func New() (Env, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	return &env{
		db: wrapper.NewQuerier(db),
	}, nil
}

// Mock impl
func Mock(db wrapper.Querier) Env {
	return &mock{
		db: db,
	}
}

type mock struct {
	db wrapper.Querier
}

func (e *mock) DB() wrapper.Querier {
	return e.db
}

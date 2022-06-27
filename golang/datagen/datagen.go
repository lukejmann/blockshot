package datagen

import (
	"github.com/lukejmann/blockshot/golang/env"
	"github.com/machinebox/graphql"
)

type Datagen struct {
	env       env.Env
	gqlClient *graphql.Client
}

func New(e env.Env, zoraEndpoint string) (*Datagen, error) {
	gqlClient := graphql.NewClient(zoraEndpoint)
	return &Datagen{
		env:       e,
		gqlClient: gqlClient,
	}, nil
}

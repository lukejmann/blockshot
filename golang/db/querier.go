// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
)

type Querier interface {
	GetFlawedMintsForBlock(ctx context.Context, blockNum int32) ([]Mint, error)
	GetHighestBlock(ctx context.Context) (interface{}, error)
	GetMintsForBlock(ctx context.Context, blockNum int32) ([]Mint, error)
	InsertMintWithImageData(ctx context.Context, arg InsertMintWithImageDataParams) error
	InsertMintWithImageURL(ctx context.Context, arg InsertMintWithImageURLParams) error
}

var _ Querier = (*Queries)(nil)

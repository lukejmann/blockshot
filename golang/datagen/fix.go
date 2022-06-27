package datagen

import (
	"context"
	"fmt"
)

type TD struct {
	tokenURL  string
	imageURL  string
	imageData int
}

func (dg *Datagen) FixMissingMetadataFromBlock(block int) error {
	fmt.Println("Fixing missing metadata from block", block)
	_, err := dg.env.DB().GetFlawedMintsForBlock(context.Background(), int32(block))
	if err != nil {
		return err
	}
	// datas := make([]TD, 0)
	// for _, mint := range mints {
	// 	datas = append(datas, TD{
	// 		tokenURL:  mint.TokenUrl.String,
	// 		imageURL:  mint.ImageUrl.String,
	// 		imageData: len(mint.ImageData.String),
	// 	})
	// }
	// spew.Dump(datas)
	return nil
}

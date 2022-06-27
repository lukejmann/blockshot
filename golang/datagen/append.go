package datagen

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lukejmann/blockshot/golang/db"
	"github.com/machinebox/graphql"
)

type Image struct {
	Url string `json:"url"`
}

type Token struct {
	CollectionAddress string `json:"collectionAddress"`
	TokenID           string `json:"tokenId"`
	TokenURL          string `json:"tokenUrl"`
	Image             Image  `json:"image"`
	Name              string `json:"name"`
}

type MintNode struct {
	Mint struct {
		TransactionInfo struct {
			BlockNumber int `json:"blockNumber"`
		} `json:"transactionInfo"`
	} `json:"mint"`
	Token Token `json:"token"`
}

type MintsResponse struct {
	Mints struct {
		Nodes []MintNode `json:"nodes"`
	} `json:"mints"`
}

func (dg *Datagen) FetchMintsAtBlock(block int) error {
	graphqlRequest := graphql.NewRequest(`
	query MyQuery {
		mints(filter: {}, networks: {}, sort: {sortKey: TIME, sortDirection: DESC}, pagination: {limit: 500}) {
		  nodes {
			mint {
			  transactionInfo {
				blockNumber
			  }
			}
			token {
			  collectionAddress
			  tokenId
			  tokenUrl
			  image {
				size
				url
			  }
			  name
			}
		  }
		}
	  }	  
    `)
	var mintsResp MintsResponse
	if err := dg.gqlClient.Run(context.Background(), graphqlRequest, &mintsResp); err != nil {
		return err
	}

	var tokensForBlock []MintNode
	for _, mint := range mintsResp.Mints.Nodes {
		if mint.Mint.TransactionInfo.BlockNumber == block {
			tokensForBlock = append(tokensForBlock, mint)
		}
	}

	fmt.Printf("Tokens for block %v: %v\n", block, len(tokensForBlock))
	fmt.Printf("Tokens fetched %v: %v\n", block, len(mintsResp.Mints.Nodes))
	if err := dg.AppendMints(mintsResp.Mints.Nodes); err != nil {
		fmt.Printf("Error appending mints: %v\n", err)
	}

	return nil
}

type TokenData struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (dg *Datagen) AppendMints(mints []MintNode) error {
	var errorCount int
	var urlDataCount int
	var imageDataCount int
	var standardCount int
	for _, mint := range mints {
		if strings.Contains(mint.Token.TokenURL, "data:application/json;base64") && len(strings.Split(mint.Token.TokenURL, "data:application/json;base64,")) > 1 {
			urlDataCount++
			base64String := strings.Split(mint.Token.TokenURL, "data:application/json;base64,")[1]
			bytes, err := b64.StdEncoding.DecodeString(base64String)
			if err != nil {
				fmt.Println("Failed to Decode str", err)
				continue
			}
			var tokenData TokenData
			if err := json.Unmarshal(bytes, &tokenData); err != nil {
				errorCount++
				fmt.Printf("Error unmarshalling token data for %v: %v\n", base64String, err)
				continue
			}
			err = dg.env.DB().InsertMintWithImageData(context.Background(), db.InsertMintWithImageDataParams{
				BlockNum:          int32(mint.Mint.TransactionInfo.BlockNumber),
				ImageData:         tokenData.Image,
				CollectionAddress: mint.Token.CollectionAddress,
				TokenID:           mint.Token.TokenID,
				TokenName:         tokenData.Name,
			})
			if err != nil {
				fmt.Printf("Error inserting mint with token url data: %v %v\n", err, mint)
				errorCount++
			}
		} else if strings.Contains(mint.Token.Image.Url, "data:image/svg+xml;base64") && len(strings.Split(mint.Token.TokenURL, "data:image/svg+xml;base64,")) > 1 {
			imageDataCount++
			base64String := strings.Split(mint.Token.TokenURL, "data:image/svg+xml;base64,")[1]
			bytes, err := b64.StdEncoding.DecodeString(base64String)
			if err != nil {
				fmt.Println("Failed to Decode str", err)
				continue
			}
			err = dg.env.DB().InsertMintWithImageData(context.Background(), db.InsertMintWithImageDataParams{
				BlockNum:          int32(mint.Mint.TransactionInfo.BlockNumber),
				ImageData:         string(bytes),
				CollectionAddress: mint.Token.CollectionAddress,
				TokenID:           mint.Token.TokenID,
				TokenName:         mint.Token.Name,
			})
			if err != nil {
				fmt.Printf("Error inserting mint with image data: %v %v\n", err, mint)
				errorCount++
			}
		} else {
			standardCount++
			err := dg.env.DB().InsertMintWithImageURL(context.Background(), db.InsertMintWithImageURLParams{
				BlockNum:          int32(mint.Mint.TransactionInfo.BlockNumber),
				ImageUrl:          mint.Token.Image.Url,
				CollectionAddress: mint.Token.CollectionAddress,
				TokenID:           mint.Token.TokenID,
				TokenUrl:          mint.Token.TokenURL,
				TokenName:         mint.Token.Name,
			})
			if err != nil {
				fmt.Printf("Error inserting mint with image url: %v %v\n", err, mint)
				errorCount++
			}
		}

	}
	fmt.Printf("Inserted %v mints\n %v errors\n %v url data\n %v image data\n %v standard\n", len(mints), errorCount, urlDataCount, imageDataCount, standardCount)
	return nil
}

package main

import (
	"io"
	"rewe/reweapi"

	"github.com/urfave/cli/v2"
)

func categoriesCommand(output io.Writer, baseURLFlag *cli.StringFlag) *cli.Command {
	productFlag := &cli.StringFlag{
		Name:     "product",
		Required: true,
	}

	jsonFlag := &cli.BoolFlag{
		Name: "json",
	}

	return &cli.Command{
		Name:  "categories",
		Usage: "fetch categories for a product",
		Flags: []cli.Flag{
			productFlag,
			jsonFlag,
		},
		Action: func(c *cli.Context) error {
			baseURL := c.String(baseURLFlag.Name)
			useJSON := c.Bool(jsonFlag.Name)
			product := c.String(productFlag.Name)

			fetcher := reweapi.CategoryFetcher{
				ReweClient:       reweapi.ReweClientImpl{BaseURL: baseURL},
				SearchPageParser: reweapi.SearchPageParserImpl{},
			}

			categories, err := fetcher.Fetch(product)
			if err != nil {
				return err
			}

			err = writeCategoryInfo(output, categories, useJSON)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

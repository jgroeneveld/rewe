package main

import (
	"github.com/urfave/cli/v2"
	"io"
	"rewe/reweapi"
)

func categoriesCommand(output io.Writer, baseUrlFlag *cli.StringFlag) *cli.Command {
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
			baseUrl := c.String(baseUrlFlag.Name)
			useJson := c.Bool(jsonFlag.Name)
			product := c.String(productFlag.Name)

			fetcher := reweapi.CategoriesFetcher{
				ReweClient:       reweapi.ReweClientImpl{BaseUrl: baseUrl},
				SearchPageParser: reweapi.SearchPageParserImpl{},
			}

			categories, err := fetcher.Fetch(product)
			if err != nil {
				return err
			}

			err = writeCategories(output, categories, useJson)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

package main

import (
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"rewe/reweapi"
)

func main() {
	app := NewApp(os.Stdout)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func NewApp(output io.Writer) *cli.App {
	categoriesCommand := &cli.Command{
		Name:  "categories",
		Usage: "fetch categories for a product",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "product",
				Required: true,
			},
			&cli.BoolFlag{
				Name: "json",
			},
		},
		Action: func(c *cli.Context) error {
			baseUrl := c.String("base-url")
			useJson := c.Bool("json")
			product := c.String("product")

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

	app := &cli.App{
		Name:  "rewe",
		Usage: "fetch categories for products of rewes online shop",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "base-url",
			},
		},
		Commands: []*cli.Command{
			categoriesCommand,
		},
	}
	return app
}

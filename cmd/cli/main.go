package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"rewe/reweapi"
)

func main() {
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
			useJson := c.Bool("json")

			product := c.String("product")

			fetcher := reweapi.NewCategoriesFetcher()

			categories, err := fetcher.Fetch(product)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = writeCategories(categories, useJson)
			if err != nil {
				log.Fatal(err.Error())
			}

			return nil
		},
	}

	app := &cli.App{
		Name:  "rewe",
		Usage: "fetch categories for products of rewes online shop",
		Commands: []*cli.Command{
			categoriesCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

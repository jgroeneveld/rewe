package main

import (
	"fmt"
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
		},
		Action: func(c *cli.Context) error {
			product := c.String("product")

			fetcher := reweapi.NewCategoriesFetcher()

			categories, err := fetcher.Fetch(product)
			if err != nil {
				log.Fatal(err.Error())
			}

			for _, c := range categories {
				fmt.Printf("%q\n", c)
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

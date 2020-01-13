package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
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
	// global flags
	var baseUrlFlag = &cli.StringFlag{
		Name: "base-url",
	}

	// categories command flags
	productFlag := &cli.StringFlag{
		Name:     "product",
		Required: true,
	}

	jsonFlag := &cli.BoolFlag{
		Name: "json",
	}

	categoriesCommand := &cli.Command{
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

	app := &cli.App{
		Name:  "rewe",
		Usage: "fetch categories for products of rewes online shop",
		Flags: []cli.Flag{
			baseUrlFlag,
		},
		Commands: []*cli.Command{
			categoriesCommand,
		},
	}
	return app
}

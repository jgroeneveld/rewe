package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"rewe/reweapi"
	"rewe/rewebill"
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

	pdfCommand := &cli.Command{
		Name:  "pdf",
		Usage: "read pdf",
		Flags: []cli.Flag{

		},
		Action: func(c *cli.Context) error {
			f, err := os.Open("fixtures/rechnung.pdf")
			if err != nil {
				return err
			}
			defer f.Close()

			pdf, err := rewebill.ReadPdf(f)
			if err != nil {
				return err
			}

			bill, err := rewebill.Extract(pdf)
			if err != nil {
				return err
			}

			for _, position := range bill.Positions {
				fmt.Printf("%+v\n", position)
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
			pdfCommand,
		},
	}
	return app
}

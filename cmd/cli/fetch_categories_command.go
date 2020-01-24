package main

import (
	"errors"
	"io"
	"rewe/reweapi"

	"github.com/spf13/cobra"
)

func fetchCategoriesCommand(output io.Writer) *cobra.Command {
	var baseURL string
	var useJSON bool

	cmd := &cobra.Command{
		Use:   "fetch-categories [product-query]",
		Short: "fetch categories for a product",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing product")
			}

			product := args[0]

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

	cmd.Flags().StringVar(&baseURL, "base-url", "", "set to overwrite the base-url of the rewe site")
	cmd.Flags().BoolVar(&useJSON, "json", false, "use json as output")

	return cmd
}

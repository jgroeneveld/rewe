package main

import (
	"errors"
	"io"
	"os"
	"rewe"
	"rewe/reweapi"
	"rewe/rewebill"

	"github.com/spf13/cobra"
)

func billCommand(output io.Writer) *cobra.Command {
	var baseURL string

	var cmd = &cobra.Command{
		Use:   "bill",
		Short: "fetch categories for all products in the bill",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing file to read")
			}

			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer f.Close()

			fetcher := reweapi.CategoryFetcher{
				ReweClient:       reweapi.ReweClientImpl{BaseURL: baseURL},
				SearchPageParser: reweapi.SearchPageParserImpl{},
			}

			augmentedBill, err := rewe.AugmentBill(f, rewebill.Reader, fetcher)
			if err != nil {
				return err
			}

			err = writeAugmentedBill(output, augmentedBill)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&baseURL, "base-url", "", "set to overwrite the base-url of the rewe site")

	return cmd
}

package main

import (
	"io"

	"github.com/spf13/cobra"
)

func rootCommand(output io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rewe",
		Short: "fetch categories for products of rewes online shop",
	}

	cmd.AddCommand(billCommand(output))
	cmd.AddCommand(readBillCommand(output))
	cmd.AddCommand(fetchCategoriesCommand(output))

	return cmd
}

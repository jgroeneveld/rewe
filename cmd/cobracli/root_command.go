package main

import (
	"github.com/spf13/cobra"
	"io"
)

func rootCommand(output io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rewe",
		Short: "fetch categories for products of rewes online shop",
	}

	cmd.AddCommand(readBillCommand(output))
	cmd.AddCommand(categoriesCommand(output))

	return cmd
}


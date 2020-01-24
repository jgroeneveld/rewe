package main

import (
	"errors"
	"io"
	"os"
	"rewe/rewebill"

	"github.com/spf13/cobra"
)

func readBillCommand(output io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read-bill [rechnung.pdf]",
		Short: "read bill pdf",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("missing file to read")
			}

			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer f.Close()

			bill, err := rewebill.Read(f)
			if err != nil {
				return err
			}

			err = PrettyJSONWriter{}.Write(output, bill)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

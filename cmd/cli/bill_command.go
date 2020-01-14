package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"rewe/rewebill"
)

func billCommand(output io.Writer) *cli.Command {
	return &cli.Command{
		Name:      "bill",
		Usage:     "read bill pdf",
		ArgsUsage: "./Rechnung.pdf",
		Flags: []cli.Flag{
		},
		Action: func(c *cli.Context) error {
			file := c.Args().First()
			if file == "" {
				_ = cli.ShowSubcommandHelp(c)
				return errors.New("Missing file to read")
			}

			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			bill, err := rewebill.Read(f)
			if err != nil {
				return err
			}

			for _, position := range bill.Positions {
				fmt.Fprintf(output, "%+v\n", position)
			}

			return nil
		},
	}
}

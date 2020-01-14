package main

import (
	"errors"
	"io"
	"os"
	"rewe/rewebill"

	"github.com/urfave/cli/v2"
)

func billCommand(output io.Writer) *cli.Command {
	return &cli.Command{
		Name:      "bill",
		Usage:     "read bill pdf",
		ArgsUsage: "./Rechnung.pdf",
		Flags:     []cli.Flag{},
		Action: func(c *cli.Context) error {
			file := c.Args().First()
			if file == "" {
				_ = cli.ShowSubcommandHelp(c)
				return errors.New("missing file to read")
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

			err = PrettyJSONWriter{}.Write(output, bill)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

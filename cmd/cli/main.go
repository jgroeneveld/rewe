package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func main() {
	app := NewApp(os.Stdout)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func NewApp(output io.Writer) *cli.App {
	// global flags
	var baseUrlFlag = &cli.StringFlag{
		Name: "base-url",
	}

	app := &cli.App{
		Name:  "rewe",
		Usage: "fetch categories for products of rewes online shop",
		Flags: []cli.Flag{
			baseUrlFlag,
		},
		Commands: []*cli.Command{
			categoriesCommand(output, baseUrlFlag),
			billCommand(output),
		},
	}
	return app
}


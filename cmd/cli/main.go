package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := NewApp(os.Stdout)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

// NewApp creates the App
func NewApp(output io.Writer) *cli.App {
	// global flags
	var baseURLFlag = &cli.StringFlag{
		Name: "base-url",
	}

	app := &cli.App{
		Name:  "rewe",
		Usage: "fetch categories for products of rewes online shop",
		Flags: []cli.Flag{
			baseURLFlag,
		},
		Commands: []*cli.Command{
			categoriesCommand(output, baseURLFlag),
			billCommand(output),
		},
	}
	return app
}

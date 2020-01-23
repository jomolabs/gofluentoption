package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/jomolabs/gofluentoption/pkg/cmd"
	"github.com/jomolabs/gofluentoption/pkg/options"
)

var (
	version string
)

func init() {
	registerCLI()
}

func registerCLI() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%s\n", c.App.Version)
	}
}

func main() {
	opts := options.NewBindable()

	app := &cli.App{
		Name:     "gofluentoption",
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "JomoLabs",
				Email: "hi@jomolabs.io",
			},
		},
		Action: cmd.CreateAction(opts),
		Flags:  opts.Flags(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

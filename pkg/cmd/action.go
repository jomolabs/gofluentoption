package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/jomolabs/gofluentoption/pkg/options"
	"github.com/jomolabs/gofluentoption/pkg/parse"
	"github.com/jomolabs/gofluentoption/pkg/template"
	"github.com/jomolabs/gofluentoption/pkg/util"
	"github.com/jomolabs/gofluentoption/pkg/writer"
)

func CreateAction(opts *options.Options) func(context *cli.Context) error {
	return func(context *cli.Context) error {
		if context.NArg() < 1 {
			return newExitCoder("please provide a file name")
		}

		if err := util.IsFile(context.Args().Get(0)); err != nil {
			return newExitCoder(err)
		}

		writable, err := writer.New(opts, context.Args().Get(0))
		if err != nil {
			return newExitCoder(err)
		}

		findableStructures := make([]string, 0)
		if context.NArg() > 1 {
			for i := 1; i < context.NArg(); i++ {
				findableStructures = append(findableStructures, context.Args().Get(i))
			}
		}

		typeInfo, err := parse.ParseSource(context.Args().Get(0), findableStructures, opts)
		if err != nil {
			return err
		}

		err = template.Render(typeInfo, writable)
		if err != nil {
			return err
		}

		return nil
	}
}

package options

import (
	"github.com/urfave/cli/v2"
)

type Options struct {
	AllowPrivateStructures bool
	AllowPrivateFields     bool
	Output                 string
	File                   string
	Pointerize             bool
	MakeCreateMethods      bool
	IgnoreFields           string

	flags []cli.Flag
}

func (o *Options) Flags() []cli.Flag {
	return o.flags
}

func NewBindable() *Options {
	opts := &Options{}
	opts.flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "allow-private-structures",
			Aliases: []string{
				"P",
			},
			Usage:       "Allows method set generation for private structures. Private structures to parse as specified by command-line arguments will be resolved whether this flag is set or not.",
			Value:       false,
			Destination: &opts.AllowPrivateStructures,
		},
		&cli.BoolFlag{
			Name: "allow-private-fields",
			Aliases: []string{
				"F",
			},
			Usage:       "Allows method set generation for private fields.",
			Value:       false,
			Destination: &opts.AllowPrivateFields,
		},
		&cli.StringFlag{
			Name: "output",
			Aliases: []string{
				"o",
			},
			Usage:       "Output medium. One of \"console\" or \"file\". Default is \"file\". If no file is specified with the \"--file\" or \"-f\" flags in \"file\" mode, the name \"zz_generated_<FILENAME>_methods.go\" will be used.",
			Value:       "file",
			Destination: &opts.Output,
		},
		&cli.BoolFlag{
			Name: "pointerize",
			Aliases: []string{
				"p",
			},
			Usage:       "When generating methods, use a pointer to the object instead of the object itself.",
			Value:       false,
			Destination: &opts.Pointerize,
		},
		&cli.StringFlag{
			Name: "file",
			Aliases: []string{
				"f",
			},
			Usage:       "File name when \"--output\" or \"-o\" is \"file\".",
			Value:       "",
			Destination: &opts.File,
		},
		&cli.BoolFlag{
			Name: "make-create-methods",
			Aliases: []string{
				"m",
			},
			Usage:       "Create \"New\" methods for types. Two methods are created: one returning an unpopulated object; the other, with a full list of arguments from the structure itself.",
			Value:       false,
			Destination: &opts.MakeCreateMethods,
		},
		&cli.StringFlag{
			Name: "ignore-fields",
			Aliases: []string{
				"i",
			},
			Usage:       "Do not generate methods for fields with these names. Specify as a comma-separated list. Optionally, if there are structures with same-named fields, you may suppress method generation for that particular structure's field using the syntax \"<typeName>:<fieldName>\".",
			Destination: &opts.IgnoreFields,
		},
	}

	return opts
}

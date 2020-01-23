package writer

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/jomolabs/gofluentoption/pkg/options"
)

func generateFileName(fileName string) string {
	originalDir, file := path.Split(fileName)
	return fmt.Sprintf("%szz_generated_%s_methods.go", originalDir, strings.TrimSuffix(file, ".go"))
}

func New(opts *options.Options, fileName string) (io.Writer, error) {
	switch opts.Output {
	case "console":
		return os.Stdout, nil
	case "file":
		fallthrough
	default:
		outputFileName := ""
		if opts.File != "" {
			outputFileName = opts.File
		} else {
			outputFileName = generateFileName(fileName)
		}

		return os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	}
}

package util

import (
	"fmt"
	"os"
)

func IsFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("file \"%s\" is not a regular file", path)
	}

	return nil
}

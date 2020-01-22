package cmd

import (
	"errors"

	"github.com/urfave/cli/v2"
)

type exitCoder struct {
	error
}

func newExitCoder(err interface{}) cli.ExitCoder {
	switch v := err.(type) {
	case string:
		return &exitCoder{errors.New(v)}
	case error:
		return &exitCoder{v}
	default:
		panic("BUGCHECK - unrecognized type provided for newExitCoder()!")
	}
}

func (e *exitCoder) ExitCode() int {
	return 1
}

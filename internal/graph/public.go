package graph

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/google/subcommands"
)

type Config struct {
	PrintStarndart bool
	NoColor        bool
	Docs           string
	OutputType     string
	LabelFormat    string
	Clusters       bool
	ShortID        bool
	Path           string
	Out            io.Writer
}

func ExecuteGraph(ctx context.Context, config Config) error {
	err := bytes.NewBuffer([]byte{})

	cmd := &Command{
		printStandard: config.PrintStarndart,
		nocolor:       config.NoColor,
		docs:          config.Docs,
		outputType:    config.OutputType,
		labelFormat:   config.LabelFormat,
		clusters:      config.Clusters,
		shortID:       config.ShortID,
		out:           config.Out,
		err:           err,
	}

	flagSet := &flag.FlagSet{}
	flagSet.Parse([]string{config.Path})

	code := cmd.Execute(ctx, flagSet, nil)
	if err.Len() != 0 {
		return fmt.Errorf("got error while execute: %s. code: %d", err.String(), code)
	}

	if code != subcommands.ExitSuccess {
		return fmt.Errorf("execute failed. code: %d", code)
	}

	return nil
}

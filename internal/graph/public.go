package graph

import (
	"context"
	"flag"

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
}

func ExecuteGraph(ctx context.Context, config Config) subcommands.ExitStatus {
	cmd := &Command{
		printStandard: config.PrintStarndart,
		nocolor:       config.NoColor,
		docs:          config.Docs,
		outputType:    config.OutputType,
		labelFormat:   config.LabelFormat,
		clusters:      config.Clusters,
		shortID:       config.ShortID,
	}

	flagSet := &flag.FlagSet{}
	flagSet.Parse([]string{config.Path})

	return cmd.Execute(ctx, flagSet, nil)
}

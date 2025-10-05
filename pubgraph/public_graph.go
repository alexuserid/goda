package pubgraph

import (
	"context"

	"github.com/alexuserid/goda/internal/graph"
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

func DefaultConfig() Config {
	return Config{
		PrintStarndart: false,
		NoColor:        false,
		Docs:           "https://pkg.go.dev/",
		OutputType:     "dot",
		LabelFormat:    "",
		Clusters:       true,
		ShortID:        true,
		Path:           "./...:mod",
	}
}

func ExecuteGraph(ctx context.Context, config *Config) subcommands.ExitStatus {
	c := DefaultConfig()
	if config != nil {
		c = *config
	}

	return graph.ExecuteGraph(ctx, graph.Config(c))
}

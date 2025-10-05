package pubgraph

import (
	"bytes"
	"context"
	"io"

	"github.com/alexuserid/goda/internal/graph"
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
	Out            io.ReadWriter
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
		Out:            bytes.NewBuffer([]byte{}),
	}
}

func ExecuteGraph(ctx context.Context, config *Config) error {
	c := DefaultConfig()
	if config != nil {
		c = *config
	}

	return graph.ExecuteGraph(ctx, graph.Config(c))
}

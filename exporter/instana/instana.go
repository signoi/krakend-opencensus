package instana

import (
	"context"
	"errors"

	"os"

	opencensus "github.com/signoi/krakend-opencensus"
	siginstana "github.com/signoi/opencensus-exporter-instana"
)

func init() {

	opencensus.RegisterExporterFactories(func(ctx context.Context, cfg opencensus.Config) (interface{}, error) {
		return Exporter(ctx, cfg)
	})
}

func Exporter(ctx context.Context, cfg opencensus.Config) (*siginstana.Exporter, error) {

	if cfg.Exporters.Instana == nil {
		return nil, errDisabled
	}

	host := os.Getenv("INSTANA_AGENT_HOST")
	if host == "" {
		host = "localhost"
	}

	return siginstana.NewExporter(cfg.Exporters.Instana.ServiceName, host, cfg.Exporters.Instana.AgentPort), nil

}

var errDisabled = errors.New("opencensus instana exporter disabled")

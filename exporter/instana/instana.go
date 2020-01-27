package instana

import (
	"context"
	"errors"
	"fmt"

	"os"

	opencensus "github.com/signoi/krakend-opencensus"
	siginstana "github.com/signoi/opencensus-exporter-instana"
)

func init() {
	fmt.Println("initializing instana")
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
	fmt.Println("instana, exporter")
	return siginstana.NewExporter(cfg.Exporters.Instana.ServiceName, host, cfg.Exporters.Instana.AgentPort), nil

}

var errDisabled = errors.New("opencensus instana exporter disabled")

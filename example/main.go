package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/devopsfaith/krakend/transport/http/client"
	"github.com/gin-gonic/gin"

	opencensusgin "github.com/devopsfaith/krakend-opencensus/router/gin"
	opencensus "github.com/signoi/krakend-opencensus"
	"github.com/signoi/krakend-opencensus/exporter"
	_ "github.com/signoi/krakend-opencensus/exporter/influxdb"
	_ "github.com/signoi/krakend-opencensus/exporter/instana"
	_ "github.com/signoi/krakend-opencensus/exporter/jaeger"
	_ "github.com/signoi/krakend-opencensus/exporter/prometheus"
	_ "github.com/signoi/krakend-opencensus/exporter/zipkin"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case sig := <-sigs:
			log.Println("Signal intercepted:", sig)
			cancel()
		case <-ctx.Done():
		}
	}()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	fmt.Println("error", err)
	if err != nil {
		log.Fatal(err)
	}

	// Register stats and trace exporters to export the collected data.
	{
		exporter.Register(logger)

		if err := opencensus.Register(ctx, serviceConfig); err != nil {
			log.Fatal(err)
		}
	}

	bf := func(cfg *config.Backend) proxy.Proxy {
		return proxy.NewHTTPProxyWithHTTPExecutor(cfg, opencensus.HTTPRequestExecutor(client.NewHTTPClient), cfg.Decoder)
	}

	// setup the krakend router
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		ProxyFactory:   opencensus.ProxyFactory(proxy.NewDefaultFactory(opencensus.BackendFactory(bf), logger)),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: opencensusgin.New(krakendgin.EndpointHandler),
		RunServer:      router.RunServer,
	})

	routerFactory.NewWithContext(ctx).Run(serviceConfig)
}

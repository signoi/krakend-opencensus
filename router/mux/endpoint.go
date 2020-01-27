package mux

import (
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router/mux"
	opencensus "github.com/signoi/krakend-opencensus"
	"go.opencensus.io/plugin/ochttp"
)

func New(hf mux.HandlerFactory) mux.HandlerFactory {
	if !opencensus.IsRouterEnabled() {
		return hf
	}
	return func(cfg *config.EndpointConfig, p proxy.Proxy) http.HandlerFunc {
		handler := ochttp.Handler{Handler: hf(cfg, p)}
		return handler.ServeHTTP
	}
}

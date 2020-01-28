module github.com/signoi/krakend-opencensus

go 1.13

require (
	contrib.go.opencensus.io/exporter/aws v0.0.0-20190807220307-c50fb1bd7f21
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	contrib.go.opencensus.io/exporter/stackdriver v0.12.9
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/aws/aws-sdk-go v1.28.9
	github.com/devopsfaith/krakend v1.1.0
	github.com/devopsfaith/krakend-opencensus v0.0.0-20191125144520-6567e6eb9b06
	github.com/gin-gonic/gin v1.1.5-0.20170702092826-d459835d2b07
	github.com/influxdata/influxdb v1.7.9 // indirect
	github.com/kpacha/opencensus-influxdb v0.0.0-20181102202715-663e2683a27c
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.3.0
	github.com/signoi/opencensus-exporter-instana v0.0.0-20200128062352-48046c3686bf
	go.opencensus.io v0.22.2
)

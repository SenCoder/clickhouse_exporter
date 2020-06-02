package main

import (
	"flag"
	"github.com/SenCoder/clickhouse_exporter/exporter"

	"net/http"
	"net/url"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/log"
)

var (
	listeningAddress     = flag.String("telemetry.address", ":9116", "Address on which to expose metrics.")
	metricsEndpoint      = flag.String("telemetry.endpoint", "/metrics", "Path under which to expose metrics.")
	clickhouseScrapeURI  = flag.String("scrape.uri", "http://localhost:8123/", "URI to clickhouse http endpoint")
	clickhouseMetricsOly = flag.Bool("ck.metrics.only", true, "Where to expose clickhouse metrics only")
	insecure             = flag.Bool("insecure", true, "Ignore server certificate if using https")
	user                 = os.Getenv("CLICKHOUSE_USER")
	password             = os.Getenv("CLICKHOUSE_PASSWORD")
)

func main() {
	flag.Parse()

	uri, err := url.Parse(*clickhouseScrapeURI)
	if err != nil {
		log.Fatal(err)
	}
	e := exporter.NewExporter(*uri, *insecure, user, password)

	var registry prometheus.Registerer
	if *clickhouseMetricsOly {
		registry = prometheus.DefaultRegisterer
	} else {
		registry = prometheus.NewRegistry()
	}

	registry.MustRegister(e)

	log.Infof("Starting Server: %s", *listeningAddress)
	http.Handle(*metricsEndpoint, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Clickhouse Exporter</title></head>
			<body>
			<h1>Clickhouse Exporter</h1>
			<p><a href="` + *metricsEndpoint + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Fatal(http.ListenAndServe(*listeningAddress, nil))
}

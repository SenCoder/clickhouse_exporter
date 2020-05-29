module github.com/SenCoder/clickhouse_exporter

go 1.13

replace (
	github.com/Sirupsen/logrus v1.6.0 => github.com/sirupsen/logrus v1.6.0
	github.com/sirupsen/logrus v1.6.0 => github.com/Sirupsen/logrus v1.6.0
)

require (
	github.com/Sirupsen/logrus v1.6.0 // indirect
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/log v0.0.0-20151026012452-9a3136781e1f
)

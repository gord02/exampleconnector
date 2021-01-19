module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver

go 1.14

require (
	github.com/grobie/gomemcache v0.0.0-20180201122607-1f779c573665
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.18.1-0.20210119200128-ea702a906fbf
	go.uber.org/zap v1.16.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

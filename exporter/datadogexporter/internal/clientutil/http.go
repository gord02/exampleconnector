// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/clientutil"

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var (
	// JSONHeaders headers for JSON requests.
	JSONHeaders = map[string]string{
		"Content-Type":     "application/json",
		"Content-Encoding": "gzip",
	}
	// ProtobufHeaders headers for protobuf requests.
	ProtobufHeaders = map[string]string{
		"Content-Type":     "application/x-protobuf",
		"Content-Encoding": "identity",
	}
)

// NewHTTPClient returns a http.Client configured with the Agent options.
func NewHTTPClient(settings exporterhelper.TimeoutSettings, insecureSkipVerify bool) *http.Client {
	return &http.Client{
		Timeout: settings.Timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				// Disable RFC 6555 Fast Fallback ("Happy Eyeballs")
				FallbackDelay: -1 * time.Nanosecond,
			}).DialContext,
			MaxIdleConns: 100,
			// Not supported by intake
			ForceAttemptHTTP2: false,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		},
	}
}

// SetExtraHeaders appends a header map to HTTP headers.
func SetExtraHeaders(h http.Header, extras map[string]string) {
	for key, value := range extras {
		h.Set(key, value)
	}
}

func UserAgent(buildInfo component.BuildInfo) string {
	return fmt.Sprintf("%s/%s", buildInfo.Command, buildInfo.Version)
}

// SetDDHeaders sets the Datadog-specific headers
func SetDDHeaders(reqHeader http.Header, buildInfo component.BuildInfo, apiKey string) {
	reqHeader.Set("DD-Api-Key", apiKey)
	reqHeader.Set("User-Agent", UserAgent(buildInfo))
}

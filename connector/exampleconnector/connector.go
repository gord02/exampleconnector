package exampleconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// schema for connector
type connectorImp struct {
	config          Config
	metricsConsumer consumer.Metrics
	logger          *zap.Logger
}

// function to create a new connector
func newConnector(logger *zap.Logger, config component.Config) (*connectorImp, error) {
	logger.Info("Building simpleconnector connector")
	cfg := config.(*Config)

	return &connectorImp{
		config: *cfg,
		logger: logger,
	}, nil
}

// Capabilities implements the consumer interface.
func (c *connectorImp) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (c *connectorImp) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// // loop through the levels of spans
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		resourceSpan := td.ResourceSpans().At(i)

		for j := 0; j < resourceSpan.ScopeSpans().Len(); j++ {
			scopeSpan := resourceSpan.ScopeSpans().At(j)

			for k := 0; k < scopeSpan.Spans().Len(); k++ {
				span := scopeSpan.Spans().At(k)
				attrs := span.Attributes()
				mapping := attrs.AsRaw()
				for key, _ := range mapping {
					if key == c.config.AttributeName {
						// create metric only if span of trace had the specific attribute
						metrics := pmetric.NewMetrics()
						return c.metricsConsumer.ConsumeMetrics(ctx, metrics)
					}
				}
			}
		}
	}
	return nil
}

// Start implements the component.Component interface.
func (p *connectorImp) Start(ctx context.Context, _ component.Host) error {
	p.logger.Info("Starting spanmetrics connector")
	return nil
}

// Shutdown implements the component.Component interface.
func (p *connectorImp) Shutdown(context.Context) error {
	p.logger.Info("Shutting down spanmetrics connector")
	return nil
}

package timeseries

import (
	"github.com/fusakla/grabana/target/graphite"
	"github.com/fusakla/grabana/target/influxdb"
	"github.com/fusakla/grabana/target/loki"
	"github.com/fusakla/grabana/target/prometheus"
	"github.com/fusakla/grabana/target/stackdriver"
	"github.com/fusakla/sdk"
)

// WithPrometheusTarget adds a prometheus query to the graph.
func WithPrometheusTarget(query string, options ...prometheus.Option) Option {
	target := prometheus.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			RefID:          target.Ref,
			Hide:           target.Hidden,
			Expr:           target.Expr,
			IntervalFactor: target.IntervalFactor,
			Interval:       target.Interval,
			Step:           target.Step,
			LegendFormat:   target.LegendFormat,
			Instant:        target.Instant,
			Format:         target.Format,
		})

		return nil
	}
}

// WithGraphiteTarget adds a Graphite target to the table.
func WithGraphiteTarget(query string, options ...graphite.Option) Option {
	target := graphite.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithInfluxDBTarget adds an InfluxDB target to the graph.
func WithInfluxDBTarget(query string, options ...influxdb.Option) Option {
	target := influxdb.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithStackdriverTarget adds a stackdriver query to the graph.
func WithStackdriverTarget(target *stackdriver.Stackdriver) Option {
	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(target.Builder)

		return nil
	}
}

// WithLokiTarget adds a loki query to the graph.
func WithLokiTarget(query string, options ...loki.Option) Option {
	target := loki.New(query, options...)

	return func(graph *TimeSeries) error {
		graph.Builder.AddTarget(&sdk.Target{
			Hide:         target.Hidden,
			Expr:         target.Expr,
			LegendFormat: target.LegendFormat,
		})

		return nil
	}
}

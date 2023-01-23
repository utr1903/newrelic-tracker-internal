package internal

import (
	"time"

	metrics "github.com/utr1903/newrelic-tracker-internal/metrics"
)

type FlushMetric struct {
	Name       string
	Value      float64
	Timestamp  int64
	Attributes map[string]string
}

func Flush(
	mf metrics.IMetricForwarder,
	metrics []FlushMetric,
) error {

	// Add individual metrics
	for _, metric := range metrics {

		// No timestamp is given, set it to now
		if metric.Timestamp == 0 {
			metric.Timestamp = time.Now().UnixMicro()
		}

		mf.AddMetric(
			metric.Timestamp,
			metric.Name,
			"gauge",
			metric.Value,
			metric.Attributes,
		)
	}

	err := mf.Run()
	if err != nil {
		return err
	}

	return nil
}

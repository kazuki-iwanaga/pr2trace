package domain

type Metric struct{}

func (m *Metric) Type() TelemetryType {
	return TelemetryTypeMetric
}

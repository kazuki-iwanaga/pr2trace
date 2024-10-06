package domain

type TelemetryType string

const (
	TelemetryTypeSpan   TelemetryType = "Span"
	TelemetryTypeMetric TelemetryType = "Metric"
	TelemetryTypeLog    TelemetryType = "Log"
)

type Telemetry interface {
	Type() TelemetryType
}

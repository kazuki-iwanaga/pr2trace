package domain

type TelemetryExporter interface {
	Export(ts []*Telemetry) error
}

package domain

type TelemetryExporter[T Telemetry] interface {
	Export(ts []*T) error
}

type Telemetry interface {
	Span | Metric | Log
}

type Span struct{}

type Metric struct{}

type Log struct{}

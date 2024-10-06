package domain

type Log struct{}

func (l *Log) Type() TelemetryType {
	return TelemetryTypeLog
}

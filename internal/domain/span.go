package domain

type Span struct{}

func (s *Span) Type() TelemetryType {
	return TelemetryTypeSpan
}

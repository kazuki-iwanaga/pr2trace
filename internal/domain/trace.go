package domain

// Trace is a tree of spans (, metrics and logs) that represent a single operation.
type Trace struct {
	// rootSpan is the root span of the trace.
	rootSpan *Span
}

// NewTrace creates a new trace with the given root span.
func NewTrace(rootSpan *Span) *Trace {
	return &Trace{
		rootSpan: rootSpan,
	}
}

// RootSpan returns the root span of the trace.
func (t *Trace) RootSpan() *Span {
	return t.rootSpan
}

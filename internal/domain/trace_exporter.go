package domain

// TraceExporter is a gateway for trace.
type TraceExporter interface {
	// Export exports trace.
	Export(traces *Trace) error
}

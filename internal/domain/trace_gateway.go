package domain

// TraceGateway is a gateway for trace.
type TraceGateway interface {
	// Export exports trace.
	Export(trace *Trace) error
}

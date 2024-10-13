package usecase

import "go.opentelemetry.io/collector/pdata/ptrace"

type TraceGateway interface {
	Save(trace *ptrace.Traces) error
}

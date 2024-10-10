package adapter

import (
	"fmt"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type TraceOtelGateway struct{}

func NewTraceOtelGateway() *TraceOtelGateway {
	return &TraceOtelGateway{}
}

func (g *TraceOtelGateway) Export(_ *domain.Trace) error {
	// Print Debug
	fmt.Println("Exporting trace...") // nolint: forbidigo

	return nil
}

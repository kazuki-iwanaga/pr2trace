package gateway

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type TelemetryOtelExporter[T domain.Telemetry] struct{}

func (e *TelemetryOtelExporter[T]) Export(_ []*T) error {
	return nil
}

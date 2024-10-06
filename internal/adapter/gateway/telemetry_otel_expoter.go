package gateway

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type TelemetryOtelExporter struct{}

func (e *TelemetryOtelExporter) Export(_ []*domain.Telemetry) error {
	return nil
}

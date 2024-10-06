package usecase

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PrTelemetryTransformer[T domain.Telemetry] interface {
	Transform(p []*domain.Pr) ([]*T, error)
}

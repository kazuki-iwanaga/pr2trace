package usecase

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PrTelemetryTransformerUsecase[T domain.Telemetry] interface {
	Transform(p []*domain.Pr) ([]*T, error)
}

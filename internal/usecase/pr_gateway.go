package usecase

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PrGateway interface {
	Get(owner, repo string, prNumber int) (*domain.Pr, error)
}

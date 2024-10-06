package gateway

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PrGithubExtractor struct{}

func (e *PrGithubExtractor) Search(_ string) ([]*domain.Pr, error) {
	return []*domain.Pr{}, nil
}

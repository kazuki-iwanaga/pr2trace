package adapter

import (
	"fmt"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PrGhGateway struct{}

func NewPrGhGateway() *PrGhGateway {
	return &PrGhGateway{}
}

func (g *PrGhGateway) Search(_ string) ([]*domain.Pr, error) {
	fmt.Println("Searching PRs...") // nolint: forbidigo

	return nil, nil
}

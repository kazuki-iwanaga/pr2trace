package adapter

import (
	"fmt"

	"github.com/kazuki-iwanaga/pr2trace/internal/usecase"
)

type etlPresenter struct{}

func NewEtlPresenter() usecase.EtlPresenter {
	return &etlPresenter{}
}

func (p *etlPresenter) Output() *usecase.EtlOutput {
	fmt.Println("Outputting from presenter...") // nolint: forbidigo

	return &usecase.EtlOutput{}
}

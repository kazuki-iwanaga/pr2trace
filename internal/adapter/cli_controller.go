package adapter

import "github.com/kazuki-iwanaga/pr2trace/internal/usecase"

type EtlController struct {
	uc usecase.EtlUsecase
}

func NewEtlController(uc usecase.EtlUsecase) *EtlController {
	return &EtlController{
		uc: uc,
	}
}

func (c *EtlController) Execute(query string) {
	input := &usecase.EtlInput{
		Query: query,
	}

	_, _ = c.uc.Execute(input)
}

package presenter

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
	"github.com/kazuki-iwanaga/pr2trace/internal/usecase"
)

type createPRmetrySpanFromPullRequestPresenter struct{}

func NewCreatePRmetrySpanFromPullRequestPresenter() usecase.CreatePRmetrySpanFromPullRequestPresenter {
	return createPRmetrySpanFromPullRequestPresenter{}
}

func (p createPRmetrySpanFromPullRequestPresenter) Output(
	s *domain.PRmetrySpan,
) usecase.CreatePRmetrySpanFromPullRequestOutput {
	return usecase.CreatePRmetrySpanFromPullRequestOutput{
		Name:  s.GetName(),
		Start: s.GetStart(),
		End:   s.GetEnd(),
	}
}

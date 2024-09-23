package usecase

import (
	"context"
	"time"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type (
	ICreatePRmetrySpanFromPullRequestUseCase interface {
		Execute(ctx context.Context, in CreatePRmetrySpanFromPullRequestInput) (CreatePRmetrySpanFromPullRequestOutput, error)
	}

	CreatePRmetrySpanFromPullRequestInput struct {
		Owner  string
		Repo   string
		Number int
	}

	CreatePRmetrySpanFromPullRequestPresenter interface {
		Output(s *domain.PRmetrySpan) CreatePRmetrySpanFromPullRequestOutput
	}

	CreatePRmetrySpanFromPullRequestOutput struct {
		Name  string
		Start time.Time
		End   time.Time
	}

	createPRmetrySpanFromPullRequestInteractor struct {
		pullRequestRepo domain.IPullRequestRepository
		prmetrySpanRepo domain.IPRmetrySpanRepository
		presenter       CreatePRmetrySpanFromPullRequestPresenter
	}
)

func NewCreatePRmetrySpanFromPullRequestInteractor(
	pullRequestRepo domain.IPullRequestRepository,
	prmetrySpanRepo domain.IPRmetrySpanRepository,
	presenter CreatePRmetrySpanFromPullRequestPresenter,
) ICreatePRmetrySpanFromPullRequestUseCase {
	return createPRmetrySpanFromPullRequestInteractor{
		pullRequestRepo: pullRequestRepo,
		prmetrySpanRepo: prmetrySpanRepo,
		presenter:       presenter,
	}
}

func (i createPRmetrySpanFromPullRequestInteractor) Execute(
	ctx context.Context, in CreatePRmetrySpanFromPullRequestInput,
) (CreatePRmetrySpanFromPullRequestOutput, error) {
	pr, err := i.pullRequestRepo.Get(ctx, in.Owner, in.Repo, in.Number)
	if err != nil {
		return CreatePRmetrySpanFromPullRequestOutput{}, err
	}

	span := domain.NewPRmetrySpan(
		pr.GetTitle(),
		pr.GetCreatedAt(),
		pr.GetMergedAt(),
	)

	span, err = i.prmetrySpanRepo.Save(ctx, span)
	if err != nil {
		return i.presenter.Output(&domain.PRmetrySpan{}), err
	}

	return i.presenter.Output(span), nil
}

package repository

import (
	"context"

	"github.com/google/go-github/v64/github"
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PullRequestGitHubRESTRepository struct {
	client *github.Client
}

func NewPullRequestGitHubRESTRepository(client *github.Client) *PullRequestGitHubRESTRepository {
	return &PullRequestGitHubRESTRepository{client: client}
}

func (r *PullRequestGitHubRESTRepository) Get(owner, repo string, number int) (*domain.PullRequest, error) {
	pr, _, err := r.client.PullRequests.Get(context.Background(), owner, repo, number)
	if err != nil {
		return nil, err
	}

	return domain.NewPullRequest(
		owner,
		repo,
		number,
		pr.GetTitle(),
		pr.CreatedAt.Time,
		pr.MergedAt.Time,
		nil,
	), nil
}

package repository

import (
	"context"

	"github.com/google/go-github/v64/github"
)

type PullRequestGitHubRESTRepository struct {
	client *github.Client
}

func NewPullRequestGitHubRESTRepository(client *github.Client) *PullRequestGitHubRESTRepository {
	return &PullRequestGitHubRESTRepository{client: client}
}

func (r *PullRequestGitHubRESTRepository) Get(owner, repo string, number int) (*github.PullRequest, error) {
	pr, _, err := r.client.PullRequests.Get(context.Background(), owner, repo, number)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

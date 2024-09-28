package domain

import "context"

type PullRequestRepository interface {
	Fetch(ctx context.Context, owner string, repo string, number int) (*PullRequest, error)
}

type PullRequest struct {
	owner  string
	repo   string
	number int

	title    string
	openedAt string
	mergedAt string
}

func NewPullRequest(
	owner string,
	repo string,
	number int,

	title string,
	openedAt string,
	mergedAt string,
) *PullRequest {
	return &PullRequest{
		owner:  owner,
		repo:   repo,
		number: number,

		title:    title,
		openedAt: openedAt,
		mergedAt: mergedAt,
	}
}

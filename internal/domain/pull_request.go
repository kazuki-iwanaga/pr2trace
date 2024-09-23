package domain

import (
	"context"
	"time"
)

type IPullRequestRepository interface {
	Get(ctx context.Context, owner, repo string, number int) (*PullRequest, error)
}

type PullRequest struct {
	owner  string
	repo   string
	number int

	title     string
	createdAt time.Time
	mergedAt  time.Time

	events []*PullRequestEvent
}

func NewPullRequest(
	owner, repo string, number int, title string, createdAt time.Time, mergedAt time.Time, events []*PullRequestEvent,
) *PullRequest {
	return &PullRequest{
		owner:     owner,
		repo:      repo,
		number:    number,
		title:     title,
		createdAt: createdAt,
		mergedAt:  mergedAt,
		events:    events,
	}
}

func (pr *PullRequest) GetTitle() string {
	return pr.title
}

func (pr *PullRequest) GetCreatedAt() time.Time {
	return pr.createdAt
}

func (pr *PullRequest) GetMergedAt() time.Time {
	return pr.mergedAt
}

package domain

import (
	"context"
	"time"
)

type IPullRequestRepository interface {
	Fetch(ctx context.Context, owner, repo string, number int) (*PullRequest, error)
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
	owner string,
	repo string,
	number int,

	title string,
	createdAt time.Time,
	mergedAt time.Time,

	events []*PullRequestEvent,
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

package domain

import (
	"context"
	"time"
)

type PullRequestGateway interface {
	Get(ctx context.Context, owner string, repo string, number int) (*PullRequest, error)
}

type PullRequest struct {
	owner  string
	repo   string
	number int

	title    string
	openedAt time.Time
	mergedAt time.Time

	events []*PullRequestEvent
}

func NewPullRequest(
	owner string,
	repo string,
	number int,

	title string,
	openedAt time.Time,
	mergedAt time.Time,

	events []*PullRequestEvent,
) *PullRequest {
	return &PullRequest{
		owner:  owner,
		repo:   repo,
		number: number,

		title:    title,
		openedAt: openedAt,
		mergedAt: mergedAt,

		events: events,
	}
}

package domain

import (
	"context"
	"time"
)

type IPullRequestEventRepository interface {
	GetAll(ctx context.Context, owner, repo string, number int) ([]*PullRequestEvent, error)
}

type PullRequestEvent struct {
	kind      string
	createdAt time.Time
}

func NewPullRequestEvent(kind string, createdAt time.Time) *PullRequestEvent {
	return &PullRequestEvent{kind: kind, createdAt: createdAt}
}

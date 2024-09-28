package domain

import "time"

type PullRequestEvent struct {
	eventType string
	createdAt time.Time
}

func NewPullRequestEvent(eventType string, createdAt time.Time) *PullRequestEvent {
	return &PullRequestEvent{
		eventType: eventType,
		createdAt: createdAt,
	}
}

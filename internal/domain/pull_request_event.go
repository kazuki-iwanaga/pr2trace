package domain

import "time"

type PullRequestEvent struct {
	eventType string
	timestamp time.Time
}

func NewPullRequestEvent(eventType string, timestamp time.Time) *PullRequestEvent {
	return &PullRequestEvent{
		eventType: eventType,
		timestamp: timestamp,
	}
}

func (e *PullRequestEvent) Type() string {
	return e.eventType
}

func (e *PullRequestEvent) Timestamp() time.Time {
	return e.timestamp
}

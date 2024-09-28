package domain

import (
	"time"
)

type PullRequestEventType string

const (
	FirstCommit  PullRequestEventType = "first_commit"
	Opened       PullRequestEventType = "opened"
	FirstReview  PullRequestEventType = "first_review"
	LastApproval PullRequestEventType = "last_approval"
	Merged       PullRequestEventType = "merged"
)

type PullRequestEvent struct {
	eventType PullRequestEventType
	timestamp time.Time
}

func NewPullRequestEvent(eventType PullRequestEventType, timestamp time.Time) *PullRequestEvent {
	return &PullRequestEvent{
		eventType: eventType,
		timestamp: timestamp,
	}
}

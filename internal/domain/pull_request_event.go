package domain

import "time"

type PullRequestEventType string

const (
	PullRequestEventTypeCommit  PullRequestEventType = "Commit"
	PullRequestEventTypeOpen    PullRequestEventType = "Open"
	PullRequestEventTypeReview  PullRequestEventType = "Review"
	PullRequestEventTypeApprove PullRequestEventType = "Approve"
	PullRequestEventTypeMerge   PullRequestEventType = "Merge"
)

type PullRequestEvent struct {
	eventType PullRequestEventType
	timestamp time.Time
}

func NewPullRequestEvent(eventType PullRequestEventType, timestamp time.Time) *PullRequestEvent {
	return &PullRequestEvent{eventType, timestamp}
}

func (e *PullRequestEvent) EventType() PullRequestEventType {
	return e.eventType
}

func (e *PullRequestEvent) Timestamp() time.Time {
	return e.timestamp
}

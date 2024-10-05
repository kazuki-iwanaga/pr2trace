package domain

import (
	"context"
	"sort"
)

type PullRequestGateway interface {
	Get(ctx context.Context, prid PullRequestID) (*PullRequest, error)
}

type PullRequest struct {
	id       PullRequestID
	metadata PullRequestMetadata
	events   []PullRequestEvent
}

func NewPullRequest(
	id PullRequestID,
	metadata PullRequestMetadata,
	events []PullRequestEvent,
) *PullRequest {
	sort.Slice(events, func(i, j int) bool {
		return events[i].timestamp.Before(events[j].timestamp)
	})

	return &PullRequest{id, metadata, events}
}

func (pr *PullRequest) ID() PullRequestID {
	return pr.id
}

func (pr *PullRequest) Metadata() PullRequestMetadata {
	return pr.metadata
}

func (pr *PullRequest) Events() []PullRequestEvent {
	return pr.events
}

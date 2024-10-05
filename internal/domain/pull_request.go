package domain

import (
	"context"
	"errors"
	"slices"
	"sort"
)

type PullRequestGateway interface {
	Get(ctx context.Context, prid PullRequestID) (*PullRequest, error)
}

type PullRequest struct {
	id       PullRequestID
	metadata PullRequestMetadata
	events   []*PullRequestEvent
}

func NewPullRequest(
	id PullRequestID,
	metadata PullRequestMetadata,
	events []*PullRequestEvent,
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

func (pr *PullRequest) Events() []*PullRequestEvent {
	return pr.events
}

type PullRequestEventSelectMethod string

const (
	PullRequestEventSelectMethodFirst PullRequestEventSelectMethod = "First"
	PullRequestEventSelectMethodLast  PullRequestEventSelectMethod = "Last"
)

var ErrPullRequestEventNotFound = errors.New("no pull request events found")

func (pr *PullRequest) SelectEvent(
	eventType PullRequestEventType,
	method PullRequestEventSelectMethod,
) (*PullRequestEvent, error) {
	switch method {
	case PullRequestEventSelectMethodFirst:
		for _, e := range pr.events {
			if e.EventType() == eventType {
				return e, nil
			}
		}
	case PullRequestEventSelectMethodLast:
		for _, e := range slices.Backward(pr.events) {
			if e.EventType() == eventType {
				return e, nil
			}
		}
	}

	return nil, ErrPullRequestEventNotFound
}

package domain

import (
	"context"
	"errors"
	"slices"
	"sort"
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
	sort.Slice(events, func(i, j int) bool {
		return events[i].timestamp.Before(events[j].timestamp)
	})

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

type PullRequestEventSelectMethod string

const (
	PullRequestEventSelectMethodFirst PullRequestEventSelectMethod = "First"
	PullRequestEventSelectMethodLast  PullRequestEventSelectMethod = "Last"
)

var ErrPullRequestEventNotFound = errors.New("pull request event not found")

func (p *PullRequest) SelectEvent(
	eventType PullRequestEventType,
	method PullRequestEventSelectMethod,
) (*PullRequestEvent, error) {
	switch method {
	case PullRequestEventSelectMethodFirst:
		for _, e := range p.events {
			if e.Type() == eventType {
				return e, nil
			}
		}
	case PullRequestEventSelectMethodLast:
		for _, e := range slices.Backward(p.events) {
			if e.Type() == eventType {
				return e, nil
			}
		}
	}

	return nil, ErrPullRequestEventNotFound
}

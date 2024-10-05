package domain

import (
	"errors"
	"slices"
)

type PullRequestEventSelectMethod string

const (
	PullRequestEventSelectMethodFirst PullRequestEventSelectMethod = "First"
	PullRequestEventSelectMethodLast  PullRequestEventSelectMethod = "Last"
)

var ErrPullRequestEventNotFound = errors.New("no pull request events found")

func SelectPullRequestEvent(
	events []*PullRequestEvent,
	eventType PullRequestEventType,
	method PullRequestEventSelectMethod,
) (*PullRequestEvent, error) {
	switch method {
	case PullRequestEventSelectMethodFirst:
		for _, e := range events {
			if e.EventType() == eventType {
				return e, nil
			}
		}
	case PullRequestEventSelectMethodLast:
		for _, e := range slices.Backward(events) {
			if e.EventType() == eventType {
				return e, nil
			}
		}
	}

	return nil, ErrPullRequestEventNotFound
}

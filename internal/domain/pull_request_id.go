package domain

import "errors"

type PullRequestID struct {
	owner  string
	repo   string
	number int
}

var ErrInvalidPullRequestNumber = errors.New("invalid pull request number")

func NewPullRequestID(owner, repo string, number int) (*PullRequestID, error) {
	if number < 1 {
		return nil, ErrInvalidPullRequestNumber
	}

	return &PullRequestID{owner, repo, number}, nil
}

func (prid *PullRequestID) Owner() string {
	return prid.owner
}

func (prid *PullRequestID) Repo() string {
	return prid.repo
}

func (prid *PullRequestID) Number() int {
	return prid.number
}

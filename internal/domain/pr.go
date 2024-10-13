package domain

import "time"

// Pr is an abstracted model of a Pull Request.
type Pr struct {
	// owner is the repository owner (e.g. "kazuki-iwanaga" in "kazuki-iwanaga/pr2trace").
	owner string
	// repo is the repository name (e.g. "pr2trace" in "kazuki-iwanaga/pr2trace").
	repo string
	// number is the PR number.
	number int

	// title is the PR title.
	title string

	// events is the list of events belong to the PR.
	events []*PrEvent
}

// NewPr creates a new Pr.
func NewPr(owner, repo string, number int, title string, events []*PrEvent) *Pr {
	return &Pr{
		owner:  owner,
		repo:   repo,
		number: number,
		title:  title,
		events: events,
	}
}

// AddEvent creates a new PrEvent and adds it to the PR.
func (p *Pr) AddEvent(eventType PrEventType, createdAt time.Time) *PrEvent {
	e := NewPrEvent(eventType, createdAt)
	p.events = append(p.events, e)

	return e
}

// Events returns the list of events belong to the PR.
func (p *Pr) Events() []*PrEvent {
	return p.events
}

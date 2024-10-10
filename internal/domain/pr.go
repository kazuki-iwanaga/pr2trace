package domain

// Pr is a domain model for PR.
type Pr struct {
	// title is a PR title.
	title string

	// events is a list of PR events.
	events []*PrEvent
}

// NewPr creates a new Pr.
func NewPr(title string, events []*PrEvent) *Pr {
	return &Pr{
		title:  title,
		events: events,
	}
}

// Events returns the list of PR events.
func (p *Pr) Events() []*PrEvent {
	return p.events
}

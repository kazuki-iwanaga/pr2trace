package domain

// Pr is a domain model for PR.
type Pr struct {
	// events is a list of PR events.
	events []*PrEvent
}

// NewPr creates a new Pr.
func NewPr(events []*PrEvent) *Pr {
	return &Pr{
		events: events,
	}
}

// Events returns the list of PR events.
func (p *Pr) Events() []*PrEvent {
	return p.events
}

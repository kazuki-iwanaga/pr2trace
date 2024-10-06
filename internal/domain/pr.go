package domain

type Pr struct {
	events []*PrEvent
}

func (p *Pr) Events() []*PrEvent {
	return p.events
}

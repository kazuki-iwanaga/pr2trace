package domain

type PrExtractor interface {
	Fetch(owner, repo string, number int) (*Pr, error)
	Search(query string) ([]*Pr, error)
}

type Pr struct {
	_ []*PrEvent
}

type PrEvent struct{}

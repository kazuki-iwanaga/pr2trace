package domain

// PrExtractor is a gateway for PR.
type PrExtractor interface {
	// Search searches PRs by the given query.
	Search(query string) ([]*Pr, error)
}

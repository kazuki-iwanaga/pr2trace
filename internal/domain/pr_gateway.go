package domain

// PrGateway is a gateway for PR.
type PrGateway interface {
	// Search searches PRs by the given query.
	Search(query string) ([]*Pr, error)
}

package domain

type PrExtractor interface {
	Search(query string) ([]*Pr, error)
}

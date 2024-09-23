package repository

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type PRmetrySpanRepository struct {
	collection []*domain.PRmetrySpan
}

func NewPRmetrySpanRepository() *PRmetrySpanRepository {
	return &PRmetrySpanRepository{}
}

func (r *PRmetrySpanRepository) Save(span *domain.PRmetrySpan) (*domain.PRmetrySpan, error) {
	r.collection = append(r.collection, span)
	return span, nil
}

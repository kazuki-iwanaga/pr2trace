package domain

import (
	"context"
	"time"
)

type IPRmetrySpanRepository interface {
	Save(ctx context.Context, span *PRmetrySpan) (*PRmetrySpan, error)
}

type PRmetrySpan struct {
	name  string
	start time.Time
	end   time.Time
}

func NewPRmetrySpan(name string, start, end time.Time) *PRmetrySpan {
	return &PRmetrySpan{name: name, start: start, end: end}
}

func (s *PRmetrySpan) GetName() string {
	return s.name
}

func (s *PRmetrySpan) GetStart() time.Time {
	return s.start
}

func (s *PRmetrySpan) GetEnd() time.Time {
	return s.end
}

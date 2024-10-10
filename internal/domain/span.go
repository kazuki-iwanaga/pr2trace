package domain

import (
	"time"
)

// Span represents a span in a trace.
type Span struct {
	name string

	startTime time.Time
	endTime   time.Time

	// Span can have child spans.
	childSpans []*Span
}

// NewSpan creates a new span with the given name, start time and end time.
func NewSpan(name string, startTime, endTime time.Time) (*Span, error) {
	// startTime must be before endTime.
	if err := ValidateTimeOrder(startTime, endTime); err != nil {
		return nil, err
	}

	return &Span{
		name:       name,
		startTime:  startTime,
		endTime:    endTime,
		childSpans: make([]*Span, 0),
	}, nil
}

// Name returns the name of the span.
func (s *Span) Name() string {
	return s.name
}

// StartTime returns the start time of the span.
func (s *Span) StartTime() time.Time {
	return s.startTime
}

// EndTime returns the end time of the span.
func (s *Span) EndTime() time.Time {
	return s.endTime
}

// ChildSpans returns the child spans of the span.
func (s *Span) ChildSpans() []*Span {
	return s.childSpans
}

// AddChildSpan adds a child span to the span.
func (s *Span) AddChildSpan(childSpan *Span) {
	s.childSpans = append(s.childSpans, childSpan)
}

// AddChildSpans adds child spans to the span.
func (s *Span) AddChildSpans(childSpans ...*Span) {
	s.childSpans = append(s.childSpans, childSpans...)
}

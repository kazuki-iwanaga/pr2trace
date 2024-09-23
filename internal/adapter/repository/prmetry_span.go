package repository

import (
	"context"
	"sync"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type PRmetrySpanRepository struct {
	spans []*trace.Span
	mu    sync.Mutex
}

func NewPRmetrySpanRepository() *PRmetrySpanRepository {
	return &PRmetrySpanRepository{
		spans: make([]*trace.Span, 0),
		mu:    sync.Mutex{},
	}
}

func (r *PRmetrySpanRepository) Save(ctx context.Context, s *domain.PRmetrySpan) (*domain.PRmetrySpan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tp := otel.GetTracerProvider()
	tracer := tp.Tracer("PRmetrySpanRepository")

	_, span := tracer.Start(
		ctx,
		s.GetName(),
		trace.WithTimestamp(s.GetStart()),
		trace.WithAttributes(
			attribute.String("custom.span.type", "PRmetrySpan"),
		),
	)
	defer span.End(trace.WithTimestamp(s.GetEnd()))

	r.spans = append(r.spans, &span)

	return s, nil
}

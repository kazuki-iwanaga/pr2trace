package domain

import (
	"time"
)

type PRmetry struct {
	traceID string
	spans   []PRSpan
	logs    []PRLog
	metrics []PRMetric
}

func NewPRmetry() *PRmetry {
	return &PRmetry{}
}

func (p *PRmetry) AddSpan(span PRSpan) {
	p.spans = append(p.spans, span)
}

func (p *PRmetry) AddLog(log PRLog) {
	p.logs = append(p.logs, log)
}

func (p *PRmetry) AddMetric(metric PRMetric) {
	p.metrics = append(p.metrics, metric)
}

func (p *PRmetry) Export() {
	// Export to storage
}

type PRSpan struct {
	traceID      string
	spanID       string
	parendSpanID string
	name         string
	startTime    time.Time
	endTime      time.Time
}

type PRLog struct {
	traceID   string
	spanID    string
	message   string
	timestamp time.Time
}

type PRMetric struct {
	traceID   string
	spanID    string
	name      string
	value     float64
	unit      string
	timestamp time.Time
}

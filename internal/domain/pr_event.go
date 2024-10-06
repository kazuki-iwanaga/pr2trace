package domain

import "time"

type (
	PrEventType string

	PrEvent struct {
		eventType PrEventType
		createdAt time.Time
	}
)

const (
	PrEventOpened PrEventType = "Opened"
	PrEventClosed PrEventType = "Closed"
)

func NewPrEvent(eventType PrEventType, createdAt time.Time) *PrEvent {
	return &PrEvent{
		eventType: eventType,
		createdAt: createdAt,
	}
}

func (e *PrEvent) EventType() PrEventType {
	return e.eventType
}

func (e *PrEvent) CreatedAt() time.Time {
	return e.createdAt
}

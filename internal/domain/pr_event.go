package domain

import "time"

// PrEvent is a domain model for PR event.
// 1 PR has multiple events.
type PrEvent struct {
	// eventType is a type of PR event.
	// Allowing values are PrEventTypeCommitted, PrEventTypeOpened,
	// PrEventTypeReviewed, PrEventTypeAproved, and PrEventTypeMerged.
	eventType PrEventType

	// createdAt is a time when the event was created.
	createdAt time.Time
}

// NewPrEvent creates a new PrEvent.
func NewPrEvent(eventType PrEventType, createdAt time.Time) *PrEvent {
	return &PrEvent{
		eventType: eventType,
		createdAt: createdAt,
	}
}

// EventType returns the event type.
func (e *PrEvent) EventType() PrEventType {
	return e.eventType
}

// CreatedAt returns the created time.
func (e *PrEvent) CreatedAt() time.Time {
	return e.createdAt
}

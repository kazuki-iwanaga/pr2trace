package domain

// PrEventType is a type for PR event type.
// Allowing values are "Committed", "Opened", "Reviewed", "Approved", and "Merged".
type PrEventType string

const (
	PrEventTypeCommitted PrEventType = "Committed"
	PrEventTypeOpened    PrEventType = "Opened"
	PrEventTypeReviewed  PrEventType = "Reviewed"
	PrEventTypeAproved   PrEventType = "Approved"
	PrEventTypeMerged    PrEventType = "Merged"
)

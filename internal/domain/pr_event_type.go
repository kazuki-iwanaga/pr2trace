package domain

// PrEventType is a kind of PrEvent.
// Allowed values are "Committed", "Opened", "Reviewed", "Approved", and "Merged".
type PrEventType string

const (
	PrEventTypeCommitted PrEventType = "Committed"
	PrEventTypeOpened    PrEventType = "Opened"
	PrEventTypeReviewed  PrEventType = "Reviewed"
	PrEventTypeApproved  PrEventType = "Approved"
	PrEventTypeMerged    PrEventType = "Merged"
)

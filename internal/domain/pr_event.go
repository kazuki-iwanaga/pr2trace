package domain

type (
	PrEventType string
	PrEvent     struct{}
)

const (
	PrEventOpened PrEventType = "Opened"
	PrEventClosed PrEventType = "Closed"
)

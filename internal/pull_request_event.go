package internal

type PullRequestEvent struct {
	name string
	// timestamp  time.Time
	// attributes map[string]interface{}
}

func NewPullRequestEvent(name string) *PullRequestEvent {
	return &PullRequestEvent{name: name}
}

func (e *PullRequestEvent) Name() string {
	return e.name
}

package domain

type PullRequestMetadata struct {
	title string
}

func NewPullRequestMetadata(title string) *PullRequestMetadata {
	return &PullRequestMetadata{title}
}

func (prm *PullRequestMetadata) Title() string {
	return prm.title
}

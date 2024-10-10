package domain

// MultiPrRootSpan creates a root span encompassing multiple pull requests.
// This span covers the period from the first commit to the last merge event across all PRs.
func MultiPrRootSpan(name string, prs []*Pr) (*Span, error) {
	if err := ValidateNotEmptySlice(prs); err != nil {
		return nil, err
	}

	// Find the first commit and the last merge event.
	var firstCommit, lastMerge *PrEvent

	for _, pr := range prs {
		for _, e := range pr.Events() {
			switch e.EventType() { // nolint: exhaustive // Only committed and merged events are relevant.
			case PrEventTypeCommitted:
				if firstCommit == nil || e.CreatedAt().Before(firstCommit.CreatedAt()) {
					firstCommit = e
				}
			case PrEventTypeMerged:
				if lastMerge == nil || e.CreatedAt().After(lastMerge.CreatedAt()) {
					lastMerge = e
				}
			}
		}
	}

	return NewSpan(name, firstCommit.CreatedAt(), lastMerge.CreatedAt())
}

// PrRootSpan creates a root span for the given pull request.
// This span covers the period from the first commit to the last merge event of the PR.
func PrRootSpan(name string, pr *Pr) (*Span, error) {
	// Find the first commit and the last merge event.
	var firstCommit, lastMerge *PrEvent

	for _, e := range pr.Events() {
		switch e.EventType() { // nolint: exhaustive // Only committed and merged events are relevant.
		case PrEventTypeCommitted:
			if firstCommit == nil || e.CreatedAt().Before(firstCommit.CreatedAt()) {
				firstCommit = e
			}
		case PrEventTypeMerged:
			if lastMerge == nil || e.CreatedAt().After(lastMerge.CreatedAt()) {
				lastMerge = e
			}
		}
	}

	return NewSpan(name, firstCommit.CreatedAt(), lastMerge.CreatedAt())
}

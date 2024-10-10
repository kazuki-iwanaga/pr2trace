package adapter

import (
	"context"
	"errors"
	"time"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
	"github.com/shurcooL/githubv4"
)

type PrGhExtractor struct {
	client *githubv4.Client
}

func NewPrGhExtractor(client *githubv4.Client) *PrGhExtractor {
	return &PrGhExtractor{client: client}
}

type (
	pullRequestCommit struct {
		Commit struct {
			AuthoredDate time.Time
		} `graphql:"commit"`
	}

	readyForReviewEvent struct {
		CreatedAt time.Time
	}

	pullRequestReview struct {
		CreatedAt time.Time
		State     string // [PENDING, COMMENTED, APPROVED, CHANGES_REQUESTED, DISMISSED]
	}

	issueComment struct {
		CreatedAt time.Time
	}

	timelineItem struct {
		Typename            string              `graphql:"__typename"`
		PullRequestCommit   pullRequestCommit   `graphql:"... on PullRequestCommit"`
		ReadyForReviewEvent readyForReviewEvent `graphql:"... on ReadyForReviewEvent"`
		PullRequestReview   pullRequestReview   `graphql:"... on PullRequestReview"`
		IssueComment        issueComment        `graphql:"... on IssueComment"`
	}
)

var ErrUnknownTimelineItem = errors.New("unknown timeline item")

func timelineItem2PrEvent(item *timelineItem) (*domain.PrEvent, error) {
	switch item.Typename {
	case "PullRequestCommit":
		return domain.NewPrEvent(
			domain.PrEventTypeCommitted,
			item.PullRequestCommit.Commit.AuthoredDate,
		), nil
	case "ReadyForReviewEvent":
		return domain.NewPrEvent(
			domain.PrEventTypeOpened,
			item.ReadyForReviewEvent.CreatedAt,
		), nil
	case "PullRequestReview":
		switch item.PullRequestReview.State {
		case "COMMENTED":
			return domain.NewPrEvent(
				domain.PrEventTypeReviewed,
				item.PullRequestReview.CreatedAt,
			), nil
		case "APPROVED":
			return domain.NewPrEvent(
				domain.PrEventTypeApproved,
				item.PullRequestReview.CreatedAt,
			), nil
		default:
			return nil, ErrUnknownTimelineItem
		}
	case "IssueComment":
		return domain.NewPrEvent(
			domain.PrEventTypeReviewed,
			item.IssueComment.CreatedAt,
		), nil
	default:
		return nil, ErrUnknownTimelineItem
	}
}

func (r *PrGhExtractor) Get( // nolint: funlen // This function is long because of GraphQL query.
	ctx context.Context,
	owner string,
	repo string,
	number int,
) (*domain.Pr, error) {
	variables := map[string]interface{}{
		"owner":  githubv4.String(owner),
		"repo":   githubv4.String(repo),
		"number": githubv4.Int(number), // nolint: gosec // Managed by githubv4 module.
		"cursor": (*githubv4.String)(nil),
	}

	var q struct {
		Repository struct {
			Pr struct {
				Title     string
				CreatedAt time.Time
				MergedAt  time.Time

				TimelineItems struct {
					Nodes    []timelineItem
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage bool
					}
				} `graphql:"timelineItems(first: 100, after: $cursor, itemTypes: [PULL_REQUEST_COMMIT, READY_FOR_REVIEW_EVENT, PULL_REQUEST_REVIEW, ISSUE_COMMENT])"` // nolint: lll
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	var timelineItems []timelineItem

	for {
		if err := r.client.Query(ctx, &q, variables); err != nil {
			return nil, err
		}

		timelineItems = append(timelineItems, q.Repository.Pr.TimelineItems.Nodes...)

		if !q.Repository.Pr.TimelineItems.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(q.Repository.Pr.TimelineItems.PageInfo.EndCursor)
	}

	pullRequestEvents := make([]*domain.PrEvent, 0, len(timelineItems))

	for _, item := range timelineItems {
		event, err := timelineItem2PrEvent(&item)
		if err != nil {
			continue
		}

		pullRequestEvents = append(pullRequestEvents, event)
	}

	return domain.NewPr(
		q.Repository.Pr.Title,

		pullRequestEvents,
	), nil
}

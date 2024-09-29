package gateway

import (
	"context"
	"errors"
	"time"

	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
	"github.com/shurcooL/githubv4"
)

type PullRequestGitHubGateway struct {
	client *githubv4.Client
}

func NewPullRequestGitHubGateway(client *githubv4.Client) *PullRequestGitHubGateway {
	return &PullRequestGitHubGateway{client: client}
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

func timelineItem2PullRequestEvent(item *timelineItem) (*domain.PullRequestEvent, error) {
	switch item.Typename {
	case "PullRequestCommit":
		return domain.NewPullRequestEvent(
			domain.PullRequestEventTypeCommit,
			item.PullRequestCommit.Commit.AuthoredDate,
		), nil
	case "ReadyForReviewEvent":
		return domain.NewPullRequestEvent(
			domain.PullRequestEventTypeOpen,
			item.ReadyForReviewEvent.CreatedAt,
		), nil
	case "PullRequestReview":
		switch item.PullRequestReview.State {
		case "COMMENTED":
			return domain.NewPullRequestEvent(
				domain.PullRequestEventTypeReview,
				item.PullRequestReview.CreatedAt,
			), nil
		case "APPROVED":
			return domain.NewPullRequestEvent(
				domain.PullRequestEventTypeApprove,
				item.PullRequestReview.CreatedAt,
			), nil
		default:
			return nil, ErrUnknownTimelineItem
		}
	case "IssueComment":
		return domain.NewPullRequestEvent(
			domain.PullRequestEventTypeReview,
			item.IssueComment.CreatedAt,
		), nil
	default:
		return nil, ErrUnknownTimelineItem
	}
}

func (r *PullRequestGitHubGateway) Get( // nolint: funlen // This function is long because of GraphQL query.
	ctx context.Context,
	owner string,
	repo string,
	number int,
) (*domain.PullRequest, error) {
	variables := map[string]interface{}{
		"owner":  githubv4.String(owner),
		"repo":   githubv4.String(repo),
		"number": githubv4.Int(number), // nolint: gosec // Managed by githubv4 module.
		"cursor": (*githubv4.String)(nil),
	}

	var q struct {
		Repository struct {
			PullRequest struct {
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

		timelineItems = append(timelineItems, q.Repository.PullRequest.TimelineItems.Nodes...)

		if !q.Repository.PullRequest.TimelineItems.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(q.Repository.PullRequest.TimelineItems.PageInfo.EndCursor)
	}

	pullRequestEvents := make([]*domain.PullRequestEvent, 0, len(timelineItems))

	for _, item := range timelineItems {
		event, err := timelineItem2PullRequestEvent(&item)
		if err != nil {
			continue
		}

		pullRequestEvents = append(pullRequestEvents, event)
	}

	return domain.NewPullRequest(
		owner,
		repo,
		number,

		q.Repository.PullRequest.Title,
		q.Repository.PullRequest.CreatedAt,
		q.Repository.PullRequest.MergedAt,

		pullRequestEvents,
	), nil
}

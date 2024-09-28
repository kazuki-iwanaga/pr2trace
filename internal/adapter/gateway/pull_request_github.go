package gateway

import (
	"context"
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

type pullRequestCommit struct {
	Commit struct {
		AuthoredDate time.Time
	} `graphql:"commit"`
}

type readyForReviewEvent struct {
	CreatedAt time.Time `graphql:"createdAt"`
}

type timelineItem struct {
	PullRequestCommit   pullRequestCommit   `graphql:"... on PullRequestCommit"`
	ReadyForReviewEvent readyForReviewEvent `graphql:"... on ReadyForReviewEvent"`
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
				} `graphql:"timelineItems(first: 100, after: $cursor, itemTypes: [PULL_REQUEST_COMMIT,READY_FOR_REVIEW_EVENT])"` // nolint: lll
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

	var pullRequestEvents []*domain.PullRequestEvent

	for _, item := range timelineItems {
		switch {
		case item.PullRequestCommit != pullRequestCommit{}: // nolint: exhaustruct // TODO
			pullRequestEvents = append(
				pullRequestEvents,
				domain.NewPullRequestEvent(
					"PullRequestCommit",
					item.PullRequestCommit.Commit.AuthoredDate,
				),
			)
		case item.ReadyForReviewEvent != readyForReviewEvent{}: // nolint: exhaustruct // TODO
			pullRequestEvents = append(
				pullRequestEvents,
				domain.NewPullRequestEvent(
					"ReadyForReview",
					item.ReadyForReviewEvent.CreatedAt,
				),
			)
		default:
			continue
		}
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

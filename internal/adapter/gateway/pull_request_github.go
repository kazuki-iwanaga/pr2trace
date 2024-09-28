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
	}

	// nolint: unused // Will be used in the future.
	type timelineItem struct {
		typename          string `graphql:"__typename"`
		pullRequestCommit struct {
			authoredDate time.Time
		} `graphql:"... on PullRequestCommit"`
		readyForReviewEvent struct {
			createdAt time.Time
		} `graphql:"... on ReadyForReviewEvent"`
	}

	var q struct {
		repository struct {
			pullRequest struct {
				title     string
				createdAt time.Time
				mergedAt  time.Time

				timelineItems struct {
					nodes    []timelineItem
					pageInfo struct {
						endCursor   githubv4.String
						hasNextPage bool
					}
				} `graphql:"timelineItems(first: 100, after: $commentsCursor, itemTypes: [PULL_REQUEST_COMMIT,READY_FOR_REVIEW_EVENT])"` // nolint: lll
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	var timelineItems []timelineItem

	for {
		if err := r.client.Query(ctx, &q, variables); err != nil {
			return nil, err
		}

		timelineItems = append(timelineItems, q.repository.pullRequest.timelineItems.nodes...)

		if !q.repository.pullRequest.timelineItems.pageInfo.hasNextPage {
			break
		}

		variables["commentsCursor"] = githubv4.NewString(q.repository.pullRequest.timelineItems.pageInfo.endCursor)
	}

	var pullRequestEvents []*domain.PullRequestEvent

	for _, item := range timelineItems {
		switch item.typename {
		case "PullRequestCommit":
			pullRequestEvents = append(
				pullRequestEvents,
				domain.NewPullRequestEvent("PullRequestCommit", item.pullRequestCommit.authoredDate),
			)
		case "ReadyForReviewEvent":
			pullRequestEvents = append(
				pullRequestEvents,
				domain.NewPullRequestEvent("ReadyForReviewEvent", item.readyForReviewEvent.createdAt),
			)
		}
	}

	return domain.NewPullRequest(
		owner,
		repo,
		number,

		q.repository.pullRequest.title,
		q.repository.pullRequest.createdAt,
		q.repository.pullRequest.mergedAt,

		pullRequestEvents,
	), nil
}

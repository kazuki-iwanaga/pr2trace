package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/google/go-github/v64/github"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

var name = "github.com/kazuki-iwanaga/pr2otel"
var version = "unspecified (probably built without goreleaser)"

// nolint: exhaustruct, gochecknoglobals
var rootCmd = &cobra.Command{
	Use:     "pr2otel",
	Version: version,
	Short:   "Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.",
	Long: `Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.

[TODO] A longer description and some examples will be written here.`,
	Example: `  pr2otel --url https://github.com/kazuki-iwanaga/pr2otel/pull/7
  pr2otel --owner kazuki-iwanaga --repo pr2otel --number 7`,
	Run: func(cmd *cobra.Command, _ []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		// ...
		// Retrieve the target Pull Request information from flags
		// <--
		owner := viper.GetString("owner")
		repo := viper.GetString("repo")
		number := viper.GetInt("number")
		// -->

		client := github.NewClient(nil)
		pr2otel(ctx, client, owner, repo, number)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Specify the target Pull Request
	rootCmd.Flags().StringP("owner", "o", "", "Owner of the GitHub repository")
	rootCmd.Flags().StringP("repo", "r", "", "Name of the GitHub repository")
	rootCmd.Flags().IntP("number", "n", 0, "Number of the GitHub Pull Request")
	// rootCmd.MarkFlagRequired("owner")
	// rootCmd.MarkFlagRequired("repo")
	// rootCmd.MarkFlagRequired("number")

	// GitHub Token (e.g. Personal Access Token, GITHUB_TOKEN in GitHub Actions) to be used for API requests
	// rootCmd.Flags().StringP("github-token", "g", "",
	// 	"GitHub Token (e.g. Personal Access Token, GITHUB_TOKEN in GitHub Actions) to be used for API requests")

	// Enable OpenTelemetry for CLI (default: false)
	// This flag controls the otel instrumentation not for pr2otel function but for the CLI itself.
	rootCmd.Flags().BoolP("enable-cli-otel", "", false, "Enable OpenTelemetry for CLI (default: false)")

	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		constantCaseFlagName := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if err := viper.BindPFlag(constantCaseFlagName, f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := viper.BindEnv(f.Name, constantCaseFlagName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})
	viper.AutomaticEnv()
}

func pr2otel(ctx context.Context, client *github.Client, owner, repo string, number int) {
	// ...
	// Setup OpenTelemetry SDK
	// <--
	resource, err := resource.New(ctx,
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(fmt.Sprintf("%s/%s", owner, repo)),
			semconv.ServiceVersionKey.String(version),
		),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	traceSpanProcessor := sdktrace.NewBatchSpanProcessor(traceExporter,
		sdktrace.WithBatchTimeout(time.Second),
	)
	traceSampler := sdktrace.ParentBased(sdktrace.AlwaysSample())
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(resource),
		sdktrace.WithSpanProcessor(traceSpanProcessor),
		sdktrace.WithSampler(traceSampler),
	)
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	tracer := tracerProvider.Tracer(name)
	// -->

	// ...
	// Pull Request as Span
	// <--
	pr, _, err := client.PullRequests.Get(ctx, owner, repo, number)
	ctx, span := tracer.Start(
		ctx,
		pr.GetTitle(),
		trace.WithAttributes(githubPullRequest2OtelAttributes(pr)...),
		trace.WithTimestamp(pr.GetCreatedAt().Time),
	)
	defer span.End(trace.WithTimestamp(pr.GetMergedAt().Time))
	// -->

	// ...
	// Pull Request Timeline Events as Span Events
	// <--
	// nolint: exhaustruct, mnd
	opt := &github.ListOptions{
		PerPage: 100,
	}
	for {
		events, resp, err := client.Issues.ListIssueEvents(ctx, owner, repo, number, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, event := range events {
			fmt.Println(event.GetCreatedAt(), event.GetEvent())

			span.AddEvent(
				event.GetEvent(),
				trace.WithAttributes(githubIssueEvent2OtelAttributes(event)...),
				trace.WithTimestamp(event.GetCreatedAt().Time),
			)
		}

		// Pagination
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	// -->
}

func githubIssueEvent2OtelAttributes(t *github.IssueEvent) []attribute.KeyValue {
	// https://github.com/google/go-github/blob/master/github/issues_events.go
	return []attribute.KeyValue{
		attribute.String("url", t.GetURL()),
		attribute.String("event", t.GetEvent()),
		attribute.String("createdAt", t.GetCreatedAt().String()),
		attribute.String("actor", t.GetActor().GetLogin()),
	}
}

func githubPullRequest2OtelAttributes(t *github.PullRequest) []attribute.KeyValue {
	// https://github.com/google/go-github/blob/master/github/pulls.go
	return []attribute.KeyValue{
		attribute.String("url", t.GetURL()),
		attribute.String("title", t.GetTitle()),
		attribute.String("body", t.GetBody()),
		attribute.String("state", t.GetState()),
		attribute.String("createdAt", t.GetCreatedAt().String()),
		attribute.String("updatedAt", t.GetUpdatedAt().String()),
		attribute.String("closedAt", t.GetClosedAt().String()),
		attribute.String("mergedAt", t.GetMergedAt().String()),
		attribute.Int("additions", t.GetAdditions()),
		attribute.Int("deletions", t.GetDeletions()),
		attribute.Int("changedFiles", t.GetChangedFiles()),
		attribute.StringSlice("labels", func(l []*github.Label) []string {
			var labels []string
			for _, label := range l {
				labels = append(labels, label.GetName())
			}
			return labels
		}(t.Labels)),
	}
}

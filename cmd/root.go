package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/kazuki-iwanaga/pr2otel/internal"

	"github.com/google/go-github/v64/github"
	"github.com/spf13/cobra"
)

var version string = "unspecified"

// nolint: exhaustruct, gochecknoglobals
var rootCmd = &cobra.Command{
	Use:     "pr2otel",
	Version: version,
	Short:   "Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.",
	Long: `Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.

[TODO] A longer description and some examples will be written here.`,
	Example: `  pr2otel --url https://github.com/kazuki-iwanaga/pr2otel/pull/7
  pr2otel --owner kazuki-iwanaga --repo pr2otel --number 7`,
	Run: func(cmd *cobra.Command, args []string) {
		// https://opentelemetry.io/docs/languages/go/getting-started/#initialize-the-opentelemetry-sdk
		// Handle SIGINT (CTRL+C) gracefully.
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		// Set up OpenTelemetry.
		otelShutdown, err := internal.SetupOTelSDK(ctx)
		if err != nil {
			return
		}
		// Handle shutdown properly so nothing leaks.
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()

		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		number, _ := cmd.Flags().GetInt("number")

		client := github.NewClient(nil)

		pr, _, err := client.PullRequests.Get(ctx, owner, repo, number)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pr.GetTitle())

		opt := &github.ListOptions{
			PerPage: 100,
		}
		for {
			events, resp, err := client.Issues.ListIssueEvents(
				ctx, owner, repo, number, opt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, event := range events {
				fmt.Println(event.GetCreatedAt(), event.GetEvent())
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("owner", "o", "", "Owner of the GitHub repository")
	rootCmd.Flags().StringP("repo", "r", "", "Name of the GitHub repository")
	rootCmd.Flags().IntP("number", "n", 1, "Number of the GitHub Pull Request")

	rootCmd.MarkFlagRequired("owner")
	rootCmd.MarkFlagRequired("repo")
	rootCmd.MarkFlagRequired("number")
}

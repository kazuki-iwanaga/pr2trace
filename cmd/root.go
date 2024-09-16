package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v64/github"
	"github.com/spf13/cobra"
)

var version string = "unspecified (probably built without goreleaser)"

var rootCmd = &cobra.Command{
	Use:     "pr2otel",
	Version: version,
	Short:   "Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.",
	Long: `Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.

[TODO] A longer description and some examples will be written here.`,
	Example: `  pr2otel --url https://github.com/kazuki-iwanaga/pr2otel/pull/7
  pr2otel --owner kazuki-iwanaga --repo pr2otel --number 7`,
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		number, _ := cmd.Flags().GetInt("number")

		client := github.NewClient(nil)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

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
			}

			// Pagination
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
	// Specify the target Pull Request
	rootCmd.Flags().StringP("owner", "o", "", "Owner of the GitHub repository")
	rootCmd.Flags().StringP("repo", "r", "", "Name of the GitHub repository")
	rootCmd.Flags().IntP("number", "n", 0, "Number of the GitHub Pull Request")
	rootCmd.MarkFlagRequired("owner")
	rootCmd.MarkFlagRequired("repo")
	rootCmd.MarkFlagRequired("number")

	// GitHub Token (e.g. Personal Access Token, GITHUB_TOKEN in GitHub Actions) to be used for API requests
	// NOTE: This flag is not implemented yet.
	rootCmd.Flags().StringP("github-token", "g", "", "GitHub Token (e.g. Personal Access Token, GITHUB_TOKEN in GitHub Actions) to be used for API requests")

	// Enable OpenTelemetry for CLI (default: false)
	// This flag controls the otel instrumentation not for pr2otel function but for the CLI itself.
	rootCmd.Flags().BoolP("enable-cli-otel", "", false, "Enable OpenTelemetry for CLI (default: false)")
}

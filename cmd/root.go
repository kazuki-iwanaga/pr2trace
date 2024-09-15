package cmd

import (
	"fmt"
	"os"
	"context"

	"github.com/spf13/cobra"
	"github.com/google/go-github/v64/github"
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
		ctx := context.Background()
		client := github.NewClient(nil)

		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		number, _ := cmd.Flags().GetInt("number")

		pr, _, err := client.PullRequests.Get(ctx, owner, repo, number)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(pr.GetTitle())
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

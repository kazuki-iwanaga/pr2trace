package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// nolint: exhaustruct, gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "pr2otel",
	Short: "Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.",
	Long: `Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.

[TODO] A longer description and some examples will be written here.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("pull-request-url", "p", "", "URL of the GitHub Pull Request")
}

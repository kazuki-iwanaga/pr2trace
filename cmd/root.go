package cmd

import (
	"fmt"
	"os"

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
		fmt.Println("Hello World!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

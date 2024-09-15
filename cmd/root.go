package cmd

import (
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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

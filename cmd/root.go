package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// nolint: exhaustruct, gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "pr2otel",
	Short: "Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.",
	Long: `Convert GitHub Pull Request to OpenTelemetry-compatible telemetry.

[TODO] A longer description and some examples will be written here.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

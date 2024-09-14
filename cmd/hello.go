package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nolint: exhaustruct, gochecknoglobals
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print 'Hello World!'",
	Long: `Print 'Hello World!'
	
with some long description.`,
	// nolint: revive
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

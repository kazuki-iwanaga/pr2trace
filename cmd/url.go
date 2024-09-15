package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url <url>",
	Short: "Specify the Pull Request by URL",
	Long: `Specify the Pull Request by URL

[TODO] The details will be written here.`,
	Args:    cobra.ExactArgs(1),
	Example: "  pr2otel url https://github.com/kazuki-iwanaga/pr2otel/pull/7",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("url called:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
}

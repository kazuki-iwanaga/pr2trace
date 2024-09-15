package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Specify the Pull Request by URL",
	Long: `Specify the Pull Request by URL

[TODO] The details will be written here.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("url called")
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
}

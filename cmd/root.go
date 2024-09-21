package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pr2otel",
	Short: "Convert Pull Request(s) to Trace.",
	Long: `pr2trace is a CLI tool to convert Pull Request(s) to Trace.
For example:
  <TODO>`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is ./.pr2otel.yaml)")

	rootCmd.Flags().BoolP("enable-cli-otel", "", false,
		"Enable OpenTelemetry instrumentation for CLI commands")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pr2otel")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

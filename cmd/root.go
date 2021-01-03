package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "esinspect",
	Short: "esinspect is for inspecting Elasticsearch search result",
	Long:  `TODO: write long description`,
}

func Execute() error {
	return rootCmd.Execute()
}

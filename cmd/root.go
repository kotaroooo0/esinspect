package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "esinspect",
	Short: "esinspect is a tool for inspecting Elasticsearch search result",
}

func Execute() error {
	return rootCmd.Execute()
}

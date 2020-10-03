package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(start)
	root.AddCommand(version)
}

var root = &cobra.Command{
	Use: "assistant",
}

// Execute is entry point for the Assistant app
func Execute() error {
	return root.Execute()
}

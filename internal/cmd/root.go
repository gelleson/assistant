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

func Execute() error {
	return root.Execute()
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "version of the build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
	},
}

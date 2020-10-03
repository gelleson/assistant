package cmd

import (
	"alfred/pkg/ascii"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// Version of the builds following tags
	Version string
	// Commit of the builds following tags
	Commit string
	// BuildTime of the builds following tags
	BuildTime string
)

var version = &cobra.Command{
	Use:   "version",
	Short: "version of the build",
	Run: func(cmd *cobra.Command, args []string) {
		banner := ascii.NewBannerWithColor(ascii.BANNER_COLOR_RED)
		banner.Print("ASSISTANT", "VERSION")
		fmt.Println(fmt.Sprintf("Version: %v", Version))
		fmt.Println(fmt.Sprintf("Commit: %v", Commit))
		fmt.Println(fmt.Sprintf("Build Time: %v", BuildTime))
		fmt.Println()
	},
}

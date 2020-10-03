/*
 * MIT License
 *
 * Copyright (c) 2020 gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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

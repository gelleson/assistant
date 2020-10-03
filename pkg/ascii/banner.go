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

package ascii

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"strings"
)

// ColorName is subtype of the string to provide color name
type ColorName string

const (
	// BANNER_COLOR_GREEM is just green
	BANNER_COLOR_GREEM ColorName = "green"
	// BANNER_COLOR_RED is just red
	BANNER_COLOR_RED ColorName = "red"
	// BANNER_COLOR_BLUE is just blue
	BANNER_COLOR_BLUE ColorName = "blue"
)

// Banner struct to provide API to print cool ASCII arts
type Banner struct {
	color ColorName
}

// NewBanner constructor to create new Banner with default ColorName
func NewBanner() *Banner {
	return &Banner{}
}

// NewBannerWithColor constructor to create new Banner with ColorName
func NewBannerWithColor(color ColorName) *Banner {
	return &Banner{color: color}
}

// Print to print as ASCII
func (b Banner) Print(words ...string) {

	if b.color == "" {
		b.color = BANNER_COLOR_BLUE
	}
	fmt.Println("")
	newFigure := figure.NewColorFigure(strings.Join(words, " "), "small", string(b.color), true)
	newFigure.Print()
	fmt.Println("")
	fmt.Println("")
}

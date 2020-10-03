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

package ascii

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"strings"
)

type ColorName string

const (
	BANNER_COLOR_GREEM ColorName = "green"
	BANNER_COLOR_RED   ColorName = "red"
	BANNER_COLOR_BLUE  ColorName = "blue"
)

type Banner struct {
	color ColorName
}

func NewBanner() *Banner {
	return &Banner{}
}

func NewBannerWithColor(color ColorName) *Banner {
	return &Banner{color: color}
}

func (b Banner) HeadBanner(words ...string) {

	if b.color == "" {
		b.color = BANNER_COLOR_BLUE
	}
	fmt.Println("")
	newFigure := figure.NewColorFigure(strings.Join(words, " "), "small", string(b.color), true)
	newFigure.Print()
	fmt.Println("")
	fmt.Println("")
}

func (b Banner) SmallBanner(name string, version string) {

	if b.color == "" {
		b.color = BANNER_COLOR_BLUE
	}
	fmt.Println("")
	newFigure := figure.NewColorFigure(fmt.Sprintf("%v %v", name, version), "small", string(b.color), false)
	newFigure.Print()
	fmt.Println("")
	fmt.Println("")
}

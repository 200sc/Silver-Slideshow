package internal

import (
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	Setups = []SlideSetup{
		title,
		intro,
		background,
		core,
		end,
	}
)

type SlideSetup struct {
	Add func(int, []*static.Slide)
	Len int
}

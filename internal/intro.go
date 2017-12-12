package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	intro = SlideSetup{
		addIntro,
		3,
	}
)

func addIntro(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Simplifying Arithmetic",
		"Restoring Original Syntax",
		"Overview",
	)
}

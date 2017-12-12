package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	core = SlideSetup{
		addCore,
		7,
	}
)

func addCore(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Origin Tracking",
		"Origin Tracking",
		"Transformations",
		"Transformations",
		"Restorations",
		"Restorations",
		"Demo",
	)
}

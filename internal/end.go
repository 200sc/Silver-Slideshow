package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	end = SlideSetup{
		addEnd,
		3,
	}
)

func addEnd(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Conclusion",
		"Future Work",
		"References",
	)
}

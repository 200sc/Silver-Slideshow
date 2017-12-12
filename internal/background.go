package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	background = SlideSetup{
		addBackground,
		4,
	}
)

func addBackground(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Concrete and Abstract Syntax",
		"Concrete and Abstract Syntax",
		"Attribute Grammars",
		"Attribute Grammars",
	)
}

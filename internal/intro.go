package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	title = SlideSetup{
		addTitle,
		1,
	}
)

func addTitle(i int, sslides []*static.Slide) {

	sslides[i].Append(
		show.Title("Transforming and Restoring Silver"),
		show.TxtAt(Gnuolane44, "Patrick Stephen", .5, .5),
	)
}

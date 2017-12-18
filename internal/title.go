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
		show.Title("Transforming and Restoring Syntax Trees"),
		show.TxtAt(Gnuolane44, "Patrick Stephen", .5, .5),
		show.TxtAt(Gnuolane44, "Minnesota Extensible Language Tools Research Group", .5, .9),
	)
}

package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

func ListSlide(ss *static.Slide, header string, list ...string) {
	ss.Append(show.Header(header))
	List(ss, list...)
}

func List(ss *static.Slide, list ...string) {
	ss.Append(show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07, list...)...)
}

func CodeSlide(ss *static.Slide, header string, list ...string) {
	ss.Append(show.Header(header))
	ss.Append(show.TxtSetFrom(Gnuolane44, .25, .35, 0, .04, list...)...)
}

func DoubleListSlide(ss *static.Slide, header string, list ...string) {
	ss.Append(show.Header(header))
	left := list[:len(list)/2]
	right := list[len(list)/2:]
	ss.Append(show.TxtSetFrom(Gnuolane44, .15, .35, 0, .07, left...)...)
	ss.Append(show.TxtSetFrom(Gnuolane44, .55, .35, 0, .07, right...)...)
}

func AddHeaders(ss []*static.Slide, start int, headers ...string) {
	for i := start; i < len(ss); i++ {
		j := i - start
		if j >= len(headers) {
			return
		}
		ss[i].Append(show.Header(headers[j]))
	}
}

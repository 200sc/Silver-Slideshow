package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render"
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

func DoubleImage(ss *static.Slide, img1 string, xs1, ys1 float64, cap1,
	img2 string, xs2, ys2 float64, cap2 string) {
	totalXScale := xs1 + xs2 + .01
	leftX := (1 - totalXScale) / 2
	leftImg := show.ImageCaptionSize(img1, leftX+xs1/2, .5, xs1, ys1, Libel28, cap1)
	rightX := leftX + xs1 + .01
	fmt.Println(leftX, rightX)
	rightImg := show.ImageCaptionSize(img2, rightX+xs2/2, .5, xs2, ys2, Libel28, cap2)
	ss.Append(leftImg, rightImg)
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

func ImgDesc(ss *static.Slide, imgPath string, yScale, xScale float64, caption, desc string) {
	img := show.ImageCaptionSize(imgPath, .25, .5, xScale, yScale, Libel28, caption)
	descs := strings.Split(desc, "\n")
	for i := 1; i < len(descs); i++ {
		// remove whitespace
		descs[i] = strings.TrimSpace(descs[i])
	}
	txt := show.TxtSetFrom(Gnuolane44, .5, .35, 0, .06, descs...)
	// todo: balance positions, sizes
	ss.Append(append(txt, img)...)
}

func SlideCount(i, total int) render.Renderable {
	str := strconv.Itoa(i) + "/" + strconv.Itoa(total)
	return show.TxtAt(Gnuolane28, str, .95, .95)
}

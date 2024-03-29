package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/200sc/silver-slideshow/internal"
	"github.com/oakmound/oak"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/shape"
	"golang.org/x/image/colornames"
)

const (
	width  = 1920
	height = 1080
)

func main() {

	show.SetDims(width, height)
	show.SetTitleFont(internal.Gnuolane72)

	bz1, _ := shape.BezierCurve(
		width/15, height/5,
		width/15, height/15,
		width/5, height/15)

	bz2, _ := shape.BezierCurve(
		width-(width/15), height/5,
		width-(width/15), height/15,
		width-(width/5), height/15)

	bz3, _ := shape.BezierCurve(
		width/15, height-(height/5),
		width/15, height-(height/15),
		width/5, height-(height/15))

	bz4, _ := shape.BezierCurve(
		width-(width/15), height-(height/5),
		width-(width/15), height-(height/15),
		width-(width/5), height-(height/15))

	fill := render.NewColorBox(width-20, height-20, color.RGBA{7, 54, 130, 255})
	fill.ShiftX(10)
	fill.ShiftY(10)

	oak.Background = image.NewUniform(color.RGBA{255, 255, 255, 255})

	bkg := render.NewComposite(
		fill,
		render.BezierThickLine(bz1, colornames.White, 1),
		render.BezierThickLine(bz2, colornames.White, 1),
		render.BezierThickLine(bz3, colornames.White, 1),
		render.BezierThickLine(bz4, colornames.White, 1),
	)

	setups := internal.Setups

	total := 0

	for _, setup := range setups {
		total += setup.Len
	}

	fmt.Println("Total slides", total)

	oak.LoadingR = bkg

	sslides := static.NewSlideSet(total,
		static.Background(bkg),
	)

	nextStart := 0

	for _, setup := range setups {
		setup.Add(nextStart, sslides)
		nextStart += setup.Len
	}

	for i := range sslides {
		sslides[i].Append(internal.SlideCount(i+1, len(sslides)))
	}

	oak.SetupConfig.Screen = oak.Screen{
		Width:  width,
		Height: height,
	}
	oak.SetupConfig.FrameRate = 30
	oak.SetupConfig.DrawFrameRate = 5

	slides := make([]show.Slide, len(sslides))
	for i, s := range sslides {
		slides[i] = s
	}

	shotIndex := 0

	oak.AddCommand("allShots", func([]string) {
		for i := 0; i <= len(slides); i++ {
			rgba := oak.ScreenShot()
			f, err := os.Create("slides/slide" + strconv.Itoa(shotIndex) + ".png")
			if err != nil {
				dlog.Error(err)
				return
			}
			png.Encode(f, rgba)
			dlog.ErrorCheck(f.Close())
			shotIndex++
			time.Sleep(1000 * time.Millisecond)
			event.Trigger("KeyUpRightArrow", nil)
			time.Sleep(1000 * time.Millisecond)
		}
	})

	show.AddNumberShortcuts(len(slides))
	show.Start(slides...)
}

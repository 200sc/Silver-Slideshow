package internal

import (
	"path/filepath"

	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/render"
	"golang.org/x/image/colornames"
)

var (
	Express28  = show.FontSize(28)(show.Express)
	Gnuolane28 = show.FontSize(28)(show.Gnuolane)
	Libel28    = show.FontSize(28)(show.Libel)

	RLibel28 = show.FontColor(colornames.Blue)(Libel28)

	Express44  = show.FontSize(44)(show.Express)
	Gnuolane44 = show.FontSize(44)(show.Gnuolane)
	Libel44    = show.FontSize(44)(show.Libel)

	Express72  = show.FontSize(72)(show.Express)
	Gnuolane72 = show.FontSize(72)(show.Gnuolane)
	Libel72    = show.FontSize(72)(show.Libel)

	// todo: access os fonts?
	CMath = (&render.FontGenerator{
		File: fpFilter("cambria_math.ttf"),
	}).Generate()
)

// todo: we need to do this because some things
// haven't started in the engine yet (the engine
// doesn't know what our directories are for assets)
// Can we change this?
func fpFilter(file string) string {
	return filepath.Join("assets", "font", file)
}

package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	intro = SlideSetup{
		addIntro,
		4,
	}
)

func addIntro(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Simplifying Arithmetic",
		"Simplifying Arithmetic",
		"Restoring Original Syntax",
		"Overview",
	)
	ImgDesc(sslides[i], "tree_diagram_1.png", .3, .4,
		"1 * square(-3) as a tree",
		`Suppose we want to simplify multiplication. 
		 The input "1 * square(-3)" is shown
		 on the left. We want to make it so
		 that 1 * x is converted into x, and
		 represent that as (mul(1, x) -> x).
		 
		 This syntax is that of a "Rewrite Rule".`,
	)
	ImgDesc(sslides[i+1], "tree_diagram_2.png", .3, .4,
		"1 * square(-3), simplified",
		`Suppose we want to simplify multiplication. 
		 The input "1 * square(-3)" is shown
		 on the left. We want to make it so
		 that 1 * x is converted into x, and
		 represent that as (mul(1, x) -> x).
		 
		 This syntax is that of a "Rewrite Rule".`,
	)
	ImgDesc(sslides[i+2], "tree_diagram_3.png", .5, .5,
		`Applying (mul(1, x) -> x) and (neg(x) -> sub(0, x)) as transformations.
		 Top: Original code and tree. Bottom: Restored Code and Transformed Tree.`,
		`Suppose we want to get the original synax
		 back from the transformed tree. This requires
		 tracking transformations, and tracking when
		 changes like "square" becoming "mul" happen.
		 
		 A transformation might produce something meaningfully
		 different from the original code. For this we need to
		 be able to fabricate new code from a tree. We call this
		 Restoration.`,
	)
	List(sslides[i+3],
		"- Background",
		"- Origin Tracking",
		"- Performing Transformations",
		"- Performing Restorations",
	)
}

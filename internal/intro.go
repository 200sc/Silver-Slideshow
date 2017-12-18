package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	intro = SlideSetup{
		addIntro,
		5,
	}
)

func addIntro(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Transforming Arithmetic Expressions",
		"Transforming Arithmetic Expressions",
		"Transforming Arithmetic Expressions",
		"Restoring Original Syntax",
		"Restoring Original Syntax",
	)
	// Just mulitiplication, no square
	ImgDesc(sslides[i], "mul1example.png", .45, .45,
		"",
		`- "mul(1, x) -> x" is shown to the left, an example
		 of a Rewrite Rule
		 
		 - We refer to this rule as 'simplify' or 'simp'`,
	)

	ImgDesc(sslides[i+1], "negexample.png", .45, .45,
		"",
		`- "neg(x) -> sub(0, x)" is shown to the left
		
		 - We refer to this rule as 'expand' or 'expd'`,
	)
	// SAY
	// But in addition to that, we want to be able to
	// take the transformed tree and create a string that
	// represents it. In the transformation here, called "expd" or expand,
	// we take negations and make them into subtrasctions with 0,
	// and we'd like for this to be reflected when we get a string
	// back from this transformed tree.

	// This is also restoration, but restoring to reconstructed syntax.

	// Square vs mul
	ImgDesc(sslides[i+2], "2by2.png", .45, .45,
		"",
		`- Different code syntax can translate to the same
		tree syntax
		
		 - Doing this helps us avoid repeating logic`,
	)
	// SAY
	// However we're representing these trees in our
	// language, we've decided that multiply and square
	// eventually become the same thing, like we see here--
	// it's not important to our calculation of the math
	// that this was a square and not a mulitiply, and
	// this lets us avoid repeating code.

	// Square vs mul, we want to get the old name back
	ImgDesc(sslides[i+3], "tree_diagram_3.png", .65, .65,
		``,
		`- Restoration describes getting old syntax back 
		from our transformed trees
		
		 - Want to retain transformations in the restored output`,
	)

	// ( tree_diagram_3.png)

	// SAY
	// But while our internal representation might not
	// have multiple node types for these different
	// concrete or input types, it's still a valid
	// ask to take this and get back the text that made it.

	// So we've established that we both want to transform
	// our representations and recall back those transformations
	// to the original strings that wrote them.  We call this Restoration,
	// in this case restoring to original syntax.

	// Square(neg), we can't get the old name
	ImgDesc(sslides[i+4], "tree_diagram_4.png", .6, .6,
		``,
		`- Some transformations disable restoring original syntax
		
		 - Because negation is transformed underneath square,
		 it might be impossible to fit the output as square's child
		
		 - This demands a way to recreate code from any node`,
	)

	// SLIDE
	// Some transformations disable original syntax restoration.

	// SAY
	// But some transformations disable us from actually performing
	// restorations and maintaining original node strings. In this case,
	// we're limited from restoring the square text because it's child
	// node was converted.

	// The way that we can restore to original syntax instead of
	// new syntax involes directly copying the original syntax
	// from the input, which is impossible if something within
	// that syntax needs to be changed.

	// This is something which future research could look into
	// resolving from context, but the way we solve this now
	// involves additional rewrite rules, letting us know that
	// mul(x,y) was x * y
}

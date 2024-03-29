package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	core = SlideSetup{
		addCore,
		10,
	}
)

func addCore(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"",
		"Origin Tracking",
		"Origin Tracking",
		"Origin Tracking",
		"",
		"Transformations",
		"Transformations",
		"Restorations",
		"Restorations",
	)

	ListSlide(sslides[i], "Overview",
		"We want to:",
		"- Transform abstract syntax",
		"- Restore original concrete syntax",
		"- Create new syntax for transformed nodes when needed",
		"And do all of this easily",
	)

	BigImgDesc(sslides[i+1], "transformation_diagram_7.png", .65, .65, "",
		`- Each node in the tree tracks what 
		node created it
		 
		- Abstract nodes can follow origins
		back until they reach a concrete node
		for restoration`,
	)

	ImgDesc(sslides[i+2], "origin_tracking_applied.png", .6, .6,
		``,
		`- Attributes can be accessed on origin nodes without 
		knowing the type of the creating node`,
	)

	ImgDesc(sslides[i+3], "attribute_access.png", .6, .6,
		``,
		`- Accessing arbitrary origin attributes requires significant
		boilerplate without code generation
		
		 - We provide a much more terse syntax`,
	)

	sslides[i+4].Append(show.Title("Demo"))

	List(sslides[i+5],
		"- Transformations are themselves attributes on all tree nodes.",
		"- Without code generation, these attributes demand a",
		"    significant amount of duplicate code.",
		"- Our method generates code from minimal rewrite rules.",
	)
	DoubleImage(sslides[i+6],
		"no_code_gen_trans.png", .65, .65, "Code to apply 'simp' and 'expd' without generating code",
		"code_gen_trans.png", .25, .4, "What we provide to generate 'simp' and 'expd'",
	)
	ImgDesc(sslides[i+7], "restore_rules.png", .45, .45,
		"",
		`- Transformed nodes must use rewrite rules to 
		restore to concrete syntax
   
	- Rewrite rules to restore concrete syntax are
	declared for each grammar
	
	- Non-transformed nodes reuse existing syntax by
	referring to their origin pointers`,
	)

	BigImgDesc(sslides[i+8], "transformation_diagram_5.png", .65, .65, "",
		`- Transformed nodes must use rewrite 
		rules to restore to concrete syntax
   
	- Rewrite rules to restore concrete syntax
	are declared for each grammar
	
	- Non-transformed nodes reuse existing
	syntax by referring to their origin 
	pointers`,
	)

	sslides[i+9].Append(show.Title("Demo"))
	// document what we're going to demo
	// 1. 1 + 2 - 3, no transformations
	// 2. -3, simple negation
	// 3. 1 * 4, simple multiplication
	// 4. -(1 * sq(2)), the paper example
	// 5. 2 * 2, to ensure it doesn't always happen
	// 6. origin attributes
	// 7. optional origin attributes
}

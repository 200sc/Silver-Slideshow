package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	core = SlideSetup{
		addCore,
		8,
	}
)

func addCore(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Origin Tracking",
		"Origin Tracking",
		"Origin Tracking: Application",
		"Transformations",
		"Transformations",
		"Restorations",
		"Restorations",
	)
	List(sslides[i],
		"- Each node in the tree tracks what node created it",
		"- Distinct from parent-child relationship",
		"- Can be followed back to source code locations",
	)
	ImgDesc(sslides[i+1], "transformation_diagram_7.png", .5, .5,
		"'-(1 * square(2))', with origin pointers",
		`- Origin Tracking enables the detection of transformations.
		 
		 - Abstract nodes can follow origins back until they reach
			 a concrete node for restoration.
			 
		 - Attributes can be accessed on origin nodes without knowing
		     the type of the creating node.`,
	)
	ImgDesc(sslides[i+2], "origin_tracking_applied.png", .4, .5,
		`Example of applying Origin Tracking to use source attributes
		 Solves having to define ambiguities dependent on the creating node`,
		`- Origin Tracking enables the detection of transformations.
		 
		 - Abstract nodes can follow origins back until they reach
			 a concrete node for restoration.
			 
		 - Attributes can be accessed on origin nodes without knowing
		     the type of the creating node.`,
	)
	List(sslides[i+3],
		"- Transformations are themselves attributes on all tree nodes.",
		"- Without code generation, these attributes demand a",
		"    significant amount of duplicate code.",
		"- Our method generates code from minimal rewrite rules.",
	)
	DoubleImage(sslides[i+4],
		"no_code_gen_trans.png", .5, .5, "Code to apply 'simp' and 'expd' without generating code",
		"code_gen_trans.png", .25, .4, "What we provide to generate 'simp' and 'expd'",
	)
	ImgDesc(sslides[i+5], "restore_rules.png", .4, .4,
		"Rule definitions to restore transformed nodes in the arithmetic grammar",
		`- Transformed nodes must use rewrite rules to restore to 
		concrete syntax.
   
		- Rewrite rules to restore concrete syntax are declared 
			for each grammar. Some of these rules can be 
			automatically detected.
	
	- Non-transformed nodes reuse existing syntax by referring 
		to their origin pointers.`,
	)
	ImgDesc(sslides[i+6], "transformation_diagram_5.png", .5, .5,
		"'-(1 * square(2))', with origin pointers and restored concrete pointers",
		`- Transformed nodes must use rewrite rules to restore to 
		     concrete syntax.
		
		 - Rewrite rules to restore concrete syntax are declared 
			 for each grammar. Some of these rules can be 
			 automatically detected.
		 
		 - Non-transformed nodes reuse existing syntax by referring 
		     to their origin pointers.`,
	)
	sslides[i+7].Append(show.Title("Demo"))
	// document what we're going to demo
	// 1. 1 + 2 - 3, no transformations
	// 2. -3, simple negation
	// 3. 1 * 4, simple multiplication
	// 4. -(1 * sq(2)), the paper example
	// 5. origin attributes
	// 6. optional origin attributes
}

package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	background = SlideSetup{
		addBackground,
		7,
	}
)

func addBackground(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Concrete and Abstract Syntax",
		"Attribute Grammars",
		"Concrete Syntax",
		"Concrete Syntax",
		"Concrete Syntax",
		"Abstract Syntax",
		"Abstract Syntax",
	)

	ImgDesc(sslides[i], "transformation_diagram_6.png", .55, .55,
		"'-(1 * square(2))'",
		`- Compilers usually represent programming languages in
		 two syntax trees: concrete and abstract.
		 
		 - The concrete syntax tree (CST) is a literal representation 
		 of the input.
		 
		 - The abstract syntax tree (AST) is a semantic representation
		 of the input.`,
	)

	ImgDesc(sslides[i+1], "ag_tree.png", .35, .45,
		"An arithmetic tree annotated with integer value attributes",
		`Attribute Grammars (AGs) are an extension of 
		 Context Free Grammars. They allow nodes to be 
		 decorated with values called attributes that can
		 be passed up and down syntax trees to evaluate 
		 the semantic meaning of the tree.`,
	)

	sslides[i+2].Append(show.ImageCaptionSize("concrete_grammar_1.png", .5, .57, .77, .77,
		Libel28,
		`The concrete syntax used in this presentation. An ETF (Expression-Term-Factor) grammar.`),
	)
	sslides[i+3].Append(show.ImageCaptionSize("concrete_grammar_2.png", .5, .57, .77, .77,
		Libel28,
		`The concrete syntax used in this presentation. An ETF (Expression-Term-Factor) grammar.`),
	)
	sslides[i+4].Append(show.ImageCaptionSize("concrete_grammar.png", .5, .57, .77, .77,
		Libel28,
		`The concrete syntax used in this presentation. An ETF (Expression-Term-Factor) grammar.`),
	)
	sslides[i+5].Append(show.ImageCaptionSize("abstract_grammar_1.png", .5, .57, .73, .73,
		Libel28,
		`The abstract syntax used in this presentation.`),
	)
	sslides[i+6].Append(show.ImageCaptionSize("abstract_grammar.png", .5, .57, .73, .73,
		Libel28,
		`The abstract syntax used in this presentation.`),
	)
}

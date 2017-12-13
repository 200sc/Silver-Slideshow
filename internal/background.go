package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	background = SlideSetup{
		addBackground,
		4,
	}
)

func addBackground(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Concrete and Abstract Syntax",
		"Attribute Grammars",
		"Concrete Syntax",
		"Abstract Syntax",
	)
	ImgDesc(sslides[i], "transformation_diagram_6.png", .5, .5,
		"'-(1 * square(2))'",
		`Compilers usually represent programming languages in
		 two syntax trees: concrete and abstract.
		 
		 The concrete syntax tree (CST) is a literal representation 
		 of the input.
		 
		 The abstract syntax tree (AST) is a semantic representation
		 of the input.`,
	)
	sslides[i+2].Append(show.ImageCaptionSize("concrete_grammar.png", .5, .57, .73, .73,
		Libel28,
		`The concrete syntax used in this presentation. An ETF (Expression-Term-Factor) grammar.`),
	)
	sslides[i+3].Append(show.ImageCaptionSize("abstract_grammar.png", .5, .57, .73, .73,
		Libel28,
		`The abstract syntax used in this presentation.`),
	)
	sslides[i+1].Append(show.TxtSetFrom(Gnuolane44, .25, .35, 0, .06,
		"Attribute Grammars (AGs) are an extension of Context Free Grammars.",
		"They allow nodes to be decorated with values called attributes",
		"that can be passed up and down syntax trees to evaluate the semantic",
		"meaning of the tree.",
		"",
		"I.e. the integer value output of arithmetic. An 'add' node defines",
		"its value as the addition of its children, and so on.",
	)...)
}

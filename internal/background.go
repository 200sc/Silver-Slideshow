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

	BigImgDesc(sslides[i], "transformation_diagram_6.png", .65, .65, "",
		`- Compilers usually represent programming 
		languages in two syntax trees: concrete and
		abstract
		
		- The concrete syntax tree (CST) is a parse
		tree
		
		- The abstract syntax tree (AST) is a 
		semantic representation of the input`,
	)

	ImgDesc(sslides[i+1], "ag_tree.png", .55, .55,
		"",
		`- Attribute Grammars (AGs) are an extension of 
		Context Free Grammars 
		
		 - AGs decorate nodes with contextual attributes to 
		 derive semantic meaning from a tree`,
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

package internal

import (
	"strings"

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

	img := show.ImageCaptionSize("transformation_diagram_6.png", .30, .5, .60, .60, Libel28, "")
	descs := strings.Split(
		`- Compilers usually represent programming 
		languages in two syntax trees: concrete and
		abstract.
		
		- The concrete syntax tree (CST) is a literal 
		representation of the input.
		
		- The abstract syntax tree (AST) is a semantic
		representation of the input.`, "\n")
	for i := 1; i < len(descs); i++ {
		// remove whitespace
		descs[i] = strings.TrimSpace(descs[i])
	}
	txt := show.TxtSetFrom(Gnuolane44, .6, .35, 0, .06, descs...)
	// todo: balance positions, sizes
	sslides[i].Append(append(txt, img)...)

	ImgDesc(sslides[i+1], "ag_tree.png", .45, .45,
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

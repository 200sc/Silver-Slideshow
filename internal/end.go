package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	end = SlideSetup{
		addEnd,
		3,
	}
)

func addEnd(i int, sslides []*static.Slide) {
	AddHeaders(sslides, i,
		"Conclusion",
		"Future Work",
		"References",
	)
	sslides[i].Append(show.TxtSetFrom(Gnuolane44, .25, .35, 0, .06,
		"Implemented the following:",
		"- Origin Tracking",
		"- Attribute Propagation through origins",
		"- Applying transformations to syntax trees",
		"- Restoring concrete syntax trees from abstract syntax trees",
		"",
		"This work has potential applications in future MELT projects for",
		"specifically reconstructing languages without excess syntax.",
	)...)
	List(sslides[i+1],
		"- Split Extension into Multiple Extensions",
		"- Additional Auto-detection of Restoration Rules",
		"      (Bidrectional Transformations)",
		"- Easier Concrete Syntax",
		"- Additional Documentation",
	)
	sslides[i+2].Append(show.TxtSetFrom(Gnuolane28, .15, .35, 0, .04,
		"Origin Tracking in Attribute Grammars.",
		"Kevin Williams and Eric Van Wyk,",
		"Proceedings of the 7th International Conference on Software Language Engineering (SLE 2014), LNCS, vol. 8706, pages 282-301, Springer Verlag, 2014.",
		"",
		"Generating Attribute Grammar-based Bidirectional Transformations from Rewrite Rules.",
		"Pedro Martins, Joao Saraiva, Joao Paulo Fernandes, and Eric Van Wyk,",
		"Proceedings of the ACM SIGPLAN 2014 Workshop on Partial Evaluation and Program Manipulation (PEPM 2014), pages 63--70, ACM.",
		"",
		"Silver: an Extensible Attribute Grammar System.",
		"Eric Van Wyk, Derek Bodin, Jimin Gao, and Lijesh Krishnan,",
		"Science of Computer Programming, 75 (1-2): 39-54, Elsevier, January, 2010.",
	)...)
}

// Consider adding
// func CiteAt(ss *static.Slide, citation)

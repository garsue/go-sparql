package sparql

import (
	"strings"
)

// IRIRef https://www.w3.org/TR/rdf-sparql-query/#rIRIref
type IRIRef interface {
	Ref() string
}

// IRI https://www.w3.org/TR/rdf-sparql-query/#rIRI_REF
type IRI string

// Ref returns IRI_REF.
func (i IRI) Ref() string {
	return "<" + strings.NewReplacer(
		"<", "%3C",
		">", "%3E",
		`"`, "%22",
		" ", "%20",
		"{", "%7B",
		"}", "%7D",
		"|", "%7C",
		"\\", "%5C",
		"^", "%5E",
		"`", "%60",
	).Replace(string(i)) + ">"
}

// PrefixedName https://www.w3.org/TR/rdf-sparql-query/#rPrefixedName
type PrefixedName string

// Ref returns PrefixedName.
func (p PrefixedName) Ref() string {
	return string(p)
}

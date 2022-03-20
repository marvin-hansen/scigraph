package app

import "scigraph/src/types/kb_types"

// https://play.golang.org/p/Qg_uv_inCek
// https://freshman.tech/snippets/go/check-if-slice-contains-element/
// contains checks if a string is present in a slice
func contains(s []*kb_types.Concept, str string) bool {
	for _, v := range s {
		if v.ConceptName == str {
			return true
		}
	}
	return false
}

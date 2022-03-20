package kb_types

func NewConceptArrayArray(concepts []string) (conceptArray []*Concept) {
	for _, e := range concepts {
		c := NewConcept(e)
		conceptArray = append(conceptArray, c)
	}
	return conceptArray
}

func NewConcept(s string) *Concept {
	return &Concept{ConceptName: s}
}

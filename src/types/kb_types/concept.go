package kb_types

import "fmt"

type Concept struct {
	ConceptName string `json:"conceptName"`
}

func (s Concept) String() string {
	return fmt.Sprintf("conceptName: %v", s.ConceptName)
}

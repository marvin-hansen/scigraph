package kb_types

import (
	"fmt"
	"strconv"
)

type Publication struct {
	GUID      uint64    `json:"guid"`
	ID        string    `json:"id"`
	Doi       string    `json:"doi,omitempty"`
	Title     string    `json:"title"`
	Published string    `json:"published"`
	Updated   string    `json:"updated"`
	Comment   string    `json:"comment"`
	Summary   string    `json:"summary"`
	Author    []*Author `json:"author"`
	Link      []*Link   `json:"link,omitempty"`
	//PrimaryCategory *Category   `json:"primaryCategory,omitempty"`
	Category []*Category `json:"category,omitempty"`
	Concept  []*Concept  `json:"concept,omitempty"`
}

func (s Publication) String() string {
	return fmt.Sprintf("Publication: \n GUID %v \n ID: %v \n DOI: %v \n Title: %v \n Published : %v \n Updated: %v \n Comment: %v \n Authord: %v \n  Link: %v \n Category: %v \n   Concept: %v \n Summary: %v \n",
		s.GUID, s.ID, s.Doi, s.Title, s.Published, s.Updated, s.Comment, s.Author, s.Link,
		s.Category, s.Concept, s.Summary)
}

func (s Publication) GetGuidString() string {
	return strconv.FormatUint(s.GUID, 10)
}

func (s Publication) GetGuid() uint64 {
	return s.GUID
}

func (s Publication) Hash() uint64 {
	return s.GUID
}

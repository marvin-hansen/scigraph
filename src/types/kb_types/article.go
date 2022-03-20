package kb_types

import "fmt"

type Publication struct {
	GUID            uint64      `json:"guid"`
	ID              string      `json:"id"`
	Doi             string      `json:"doi"`
	Title           string      `json:"title"`
	Published       string      `json:"published"`
	Updated         string      `json:"updated"`
	Comment         string      `json:"comment"`
	Summary         string      `json:"summary"`
	Text            string      `json:"content"`
	Author          []*Author   `json:"author"`
	Link            []*Link     `json:"link"`
	PrimaryCategory *Category   `json:"primaryCategory,omitempty"`
	Category        []*Category `json:"category,omitempty"`
	Concept         []*Concept  `json:"concept,omitempty"`
}

func (s Publication) String() string {
	return fmt.Sprintf("Publication: \n GUID %v \n ID: %v \n DOI: %v \n Title: %v \n Published : %v \n Updated: %v \n Comment: %v \n Authord: %v \n  Link: %v \n PrimaryCategory: %v \n Category: %v \n Summary: %v \n Content: %v \n ",
		s.GUID, s.ID, s.Doi, s.Title, s.Published, s.Updated, s.Comment, s.Author, s.Link, s.PrimaryCategory,
		s.Category, s.Summary, s.Text)
}

func (s Publication) GetGuid() uint64 {
	return s.GUID
}

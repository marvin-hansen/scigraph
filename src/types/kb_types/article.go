package kb_types

import "fmt"

type Publication struct {
	GUID            uint64      `json:"guid"`
	ID              string      `json:"id"`
	Doi             string      `json:"doi"`
	Title           string      `json:"title"`
	Published       TimeStr     `json:"published"`
	Updated         TimeStr     `json:"updated"`
	Comment         string      `json:"comment"`
	Author          []*Author   `json:"author"`
	Link            []*Link     `json:"link"`
	PrimaryCategory *Category   `json:"primary_category,omitempty"`
	Category        []*Category `json:"category,omitempty"`
	Summary         *Text       `json:"summary"`
	Content         *Text       `json:"content"`
}

func (s Publication) String() string {
	return fmt.Sprintf("Publication: \n GUID %v \n ID: %v \n DOI: %v \n Title: %v \n Published : %v \n Updated: %v \n Comment: %v \n Authord: %v \n  Link: %v \n PrimaryCategory: %v \n Category: %v \n Summary: %v \n Content: %v \n ",
		s.GUID, s.ID, s.Doi, s.Title, s.Published, s.Updated, s.Comment, s.Author, s.Link, s.PrimaryCategory,
		s.Category, s.Summary, s.Content)
}

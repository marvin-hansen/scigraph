package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func newPublicationArray(entries []*arxiv.Entry) (publicationArray []*Publication) {
	for _, e := range entries {
		p := newPublication(e)
		publicationArray = append(publicationArray, p)
	}
	return publicationArray
}

func newPublication(entry *arxiv.Entry) *Publication {
	return &Publication{
		ID:              entry.ID,
		Doi:             entry.Doi,
		Title:           entry.Title,
		Link:            newLinkArray(entry.Link),
		Published:       TimeStr(entry.Published),
		Updated:         TimeStr(entry.Updated),
		Comment:         entry.Comment,
		Author:          newAuthorArray(entry.Author),
		Summary:         newText(entry.Summary.Body),
		Content:         newText(entry.Content.Body),
		PrimaryCategory: NewCategory(entry.PrimaryCategory.Term.String()),
		Category:        NewCategoryArray(entry.Category),
	}
}

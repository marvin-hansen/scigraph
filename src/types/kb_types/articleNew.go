package kb_types

import (
	"github.com/marvin-hansen/arxiv/v1"
	"scigraph/src/utils/crypto_utils"
)

func NewPublicationArray(entries []*arxiv.Entry) (publicationArray []*Publication) {
	for _, e := range entries {
		p := NewPublication(e)
		publicationArray = append(publicationArray, p)
	}
	return publicationArray
}

func NewPublication(entry *arxiv.Entry) *Publication {
	return &Publication{
		GUID:      crypto_utils.HashString(entry.ID),
		ID:        entry.ID,
		Doi:       entry.Doi,
		Title:     entry.Title,
		Comment:   entry.Comment,
		Summary:   entry.Summary.Body,
		Link:      NewLinkArray(entry.Link),
		Published: convertTimeStrToString(entry.Published),
		Updated:   convertTimeStrToString(entry.Updated),
		//
		Author:   NewAuthorArray(entry.Author),
		Category: NewCategoryArray(entry.Category),
		Concept:  nil, // set concepts during NLP processing!
	}
}

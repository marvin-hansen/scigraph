package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func NewLinkArray(links []arxiv.Link) (linkArray []*Link) {
	for _, e := range links {
		l := NewLink(e)
		linkArray = append(linkArray, l)
	}
	return linkArray
}

func NewLink(l arxiv.Link) *Link {
	return &Link{
		Rel:       l.Rel,
		Href:      l.Href,
		Type:      l.Type,
		HrefLang:  l.HrefLang,
		LinkTitle: l.Title,
		Length:    l.Length,
	}
}

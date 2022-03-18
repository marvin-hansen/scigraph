package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func newLinkArray(links []arxiv.Link) (linkArray []*Link) {
	for _, e := range links {
		l := newLink(e)
		linkArray = append(linkArray, l)
	}
	return linkArray
}

func newLink(l arxiv.Link) *Link {
	return &Link{
		Rel:      l.Rel,
		Href:     l.Href,
		Type:     l.Type,
		HrefLang: l.HrefLang,
		Title:    l.Title,
		Length:   l.Length,
	}
}

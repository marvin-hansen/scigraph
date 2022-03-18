package kb_types

import (
	"github.com/marvin-hansen/arxiv/v1"
)

func newLinkArray(links []arxiv.Link) (linkArray []*Link) {
	for _, e := range links {
		l := newLink(e)
		linkArray = append(linkArray, l) // need to dref link here to match return type to Publication type Link Array
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

type Link struct {
	Rel      string `json:"rel,omitempty"`
	Href     string `json:"href"`
	Type     string `json:"type,omitempty"`
	HrefLang string `json:"hreflang,omitempty"`
	Title    string `json:"title,omitempty"`
	Length   uint   `json:"length,omitempty"`
}

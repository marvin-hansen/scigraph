package kb_types

import (
	"fmt"
)

type Link struct {
	Rel       string `json:"rel,omitempty"`
	Href      string `json:"href"`
	Type      string `json:"type,omitempty"`
	HrefLang  string `json:"hreflang,omitempty"`
	LinkTitle string `json:"linkTitle,omitempty"`
	Length    uint   `json:"length,omitempty"`
}

func (s Link) String() string {
	return fmt.Sprintf("Link: \n Rel: %v \n Href: %v \n Type: %v \n HrefLang: %v \n LinkTitle: %v \n Length: %v ",
		s.Rel, s.Href, s.Type, s.HrefLang, s.LinkTitle, s.Length)
}

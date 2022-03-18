package kb_types

func newText(s string) *Text {
	return &Text{Body: s}
}

type Text struct {
	Body string `json:"body"`
}

package kb_types

type Text struct {
	TextBody string `json:"body"`
}

func (t Text) String() string {
	return t.TextBody
}

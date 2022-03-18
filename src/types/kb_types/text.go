package kb_types

type Text struct {
	TextBody string `json:"textBody"`
}

func (t Text) String() string {
	return t.TextBody
}

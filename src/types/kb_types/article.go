package kb_types

type Publication struct {
	ID              string      `json:"id"`
	Doi             string      `json:"doi"`
	Title           string      `json:"title"`
	Link            []*Link     `json:"link"`
	Published       TimeStr     `json:"published"`
	Updated         TimeStr     `json:"updated"`
	Comment         string      `json:"comment"`
	Author          []*Author   `json:"author"`
	Summary         *Text       `json:"summary"`
	Content         *Text       `json:"content"`
	PrimaryCategory *Category   `json:"primary_category,omitempty"`
	Category        []*Category `json:"category,omitempty"`
}

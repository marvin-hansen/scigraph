package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func NewCategoryArray(categories []*arxiv.Class) (categoryArray []*Category) {
	for _, e := range categories {
		c := NewCategory(e.Term.String())
		categoryArray = append(categoryArray, c)
	}
	return categoryArray
}

func NewCategory(s string) *Category {
	return &Category{Term: s}
}

type Category struct {
	Term string `json:"term"`
}

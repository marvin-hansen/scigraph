package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func NewAuthorArray(people []*arxiv.Person) (authorArray []*Author) {
	for _, p := range people {
		a := NewAuthor(*p)
		authorArray = append(authorArray, a)
	}
	return authorArray
}

func NewAuthor(p arxiv.Person) *Author {
	return &Author{
		Name:  p.Name,
		URI:   p.URI,
		Email: p.Email,
	}
}

package kb_types

import "github.com/marvin-hansen/arxiv/v1"

func newAuthorArray(people []*arxiv.Person) (authorArray []*Author) {
	for _, p := range people {
		a := newAuthor(*p)
		authorArray = append(authorArray, a)
	}
	return authorArray
}

func newAuthor(p arxiv.Person) *Author {
	return &Author{
		Name:  p.Name,
		URI:   p.URI,
		Email: p.Email,
	}
}

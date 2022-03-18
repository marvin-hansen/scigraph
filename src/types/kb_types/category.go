package kb_types

import "fmt"

type Category struct {
	Term string `json:"term"`
}

func (s Category) String() string {
	return fmt.Sprintf("Category: %v", s.Term)
}

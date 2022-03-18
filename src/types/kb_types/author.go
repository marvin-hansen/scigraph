package kb_types

import (
	"fmt"
)

type Author struct {
	Name  string `json:"name"`
	URI   string `json:"uri,omitempty"`
	Email string `json:"email,omitempty"`
}

func (s Author) String() string {
	return fmt.Sprintf("Author: \n Name %v \n URI: %v \n EMail: %v \n ",
		s.Name, s.URI, s.Email)
}

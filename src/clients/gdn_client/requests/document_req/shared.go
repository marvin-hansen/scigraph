package document_req

import "fmt"

type DocumentResult struct {
	Id  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
	Rev string `json:"_rev,omitempty"`
}

func (r DocumentResult) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v",
		r.Id,
		r.Key,
		r.Rev,
	)
}

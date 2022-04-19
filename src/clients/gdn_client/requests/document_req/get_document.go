package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetDocument(fabric, collectionName, key string) *RequestForGetDocument {
	return &RequestForGetDocument{
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName,
			key,
		),
	}
}

type RequestForGetDocument struct {
	path string
}

func (req *RequestForGetDocument) Path() string {
	return req.path
}

func (req *RequestForGetDocument) Method() string {
	return http.MethodGet
}

func (req *RequestForGetDocument) Query() string {
	return ""
}

func (req *RequestForGetDocument) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetDocument) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetDocument) Payload() []byte {
	return nil
}

func (req *RequestForGetDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetDocument() *ResponseForGetDocument {
	return new(ResponseForGetDocument)
}

type ResponseForGetDocument DocumentResult

func (r *ResponseForGetDocument) IsResponse() {}

func (r ResponseForGetDocument) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v, OldRev: %v",
		r.Id,
		r.Key,
		r.Rev,
		r.OldRev,
	)
}

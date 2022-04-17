package requests

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllCollections() *RequestForGetAllCollections {
	return &RequestForGetAllCollections{path: "/_api/collection?excludeSystem=true"}
}

type RequestForGetAllCollections struct {
	path string
}

func (req *RequestForGetAllCollections) Path() string {
	return req.path
}

func (req *RequestForGetAllCollections) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllCollections) Query() string {
	return ""
}

func (req *RequestForGetAllCollections) Payload() []byte {
	return nil
}

func (req *RequestForGetAllCollections) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllCollections() *ResponseForGetAllCollections {
	return new(ResponseForGetAllCollections)
}

func (r ResponseForGetAllCollections) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}

type ResponseForGetAllCollections struct {
	// @FIXME
	Field string
}

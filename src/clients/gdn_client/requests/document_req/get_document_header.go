package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetDocumentHeader(fabric string) *RequestForGetDocumentHeader {
	// @FIXME: Add correct API path
	return &RequestForGetDocumentHeader{
		path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForGetDocumentHeader struct {
	path string
}

func (req *RequestForGetDocumentHeader) Path() string {
	return req.path
}

func (req *RequestForGetDocumentHeader) Method() string {
	return http.MethodGet
}

func (req *RequestForGetDocumentHeader) Query() string {
	return ""
}

func (req *RequestForGetDocumentHeader) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetDocumentHeader) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForGetDocumentHeader) Payload() []byte {
	return nil
}

func (req *RequestForGetDocumentHeader) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetDocumentHeader() *ResponseForGetDocumentHeader {
	return new(ResponseForGetDocumentHeader)
}

type ResponseForGetDocumentHeader struct {
	// @FIXME
	Field string
}

func (r *ResponseForGetDocumentHeader) IsResponse() {}

func (r ResponseForGetDocumentHeader) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}

package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForReplaceDocument
// Replaces multiple documents in the specified collection with the ones in the body. The replaced documents are specified by the _key attributes in the body documents.
// If ignoreRevs is false, a _rev attribute in each document body must match the revision of the corresponding document in the database. Otherwise, the call fails
// In case of an error or violated precondition, an error object with the attribute error set to true and the attribute errorCode set to the error code.
//
//If the query parameter returnOld is true, for each generated document the previous revision of the document is returned under the old attribute in the result.
//If the query parameter returnNew is true, for each generated document the new document is returned under the new attribute in the result.
func NewRequestForReplaceDocument(fabric, collectionName string, jsonDocument []byte, parameters *ReplaceDocumentParameters) *RequestForReplaceDocument {
	return &RequestForReplaceDocument{
		payload: jsonDocument,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v",
			fabric, collectionName,
		),
		parameters: fmt.Sprintf("?ignoreRevs=%v&returnOld=%v&returnNew=%v&waitForSync=%v",
			parameters.ignoreRevs,
			parameters.returnOld,
			parameters.returnNew,
			parameters.waitForSync,
		),
	}
}

type RequestForReplaceDocument struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceDocument) Path() string {
	return req.path
}

func (req *RequestForReplaceDocument) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceDocument) Query() string {
	return ""
}

func (req *RequestForReplaceDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReplaceDocument() *ResponseForReplaceDocument {
	return new(ResponseForReplaceDocument)
}

type ResponseForReplaceDocument []DocumentResult

func (r *ResponseForReplaceDocument) IsResponse() {}

func (r ResponseForReplaceDocument) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()
}

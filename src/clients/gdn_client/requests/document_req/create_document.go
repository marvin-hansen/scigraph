package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForCreateDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func NewRequestForCreateDocument(fabric, collectionName string, silent bool, jsonDocument []byte, parameters *CreateDocumentParameters) *RequestForCreateDocument {

	if parameters == nil {
		parameters = GetDefaultCreateDocumentParameters()
	}

	return &RequestForCreateDocument{
		path:    fmt.Sprintf("_fabric/%v/_api/document/%v", fabric, collectionName),
		payload: jsonDocument,
		parameter: fmt.Sprintf("?returnNew=%v&returnOld=%v&silent=%v&overwrite=%v&waitForSync=%v",
			parameters.returnNew,
			parameters.returnOld,
			silent,
			parameters.overwrite,
			parameters.waitForSync,
		),
	}
}

type RequestForCreateDocument struct {
	path      string
	payload   []byte
	parameter string
}

type CreateDocumentParameters struct {
	returnNew   bool // If set to true, adds the new documents to the new attribute. False by default.
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	overwrite   bool // If set to true, the insert becomes a replace-insert. If a document with the same _key already exists the new document is not rejected with unique constraint violated but replaces the old document. False by default
	waitForSync bool // If set to true, returns only after data has been synced to disk. // False by default
}

// GetDefaultCreateDocumentParameters default values for createDocument paramters
//	returnNew   If set to true, adds the new documents to the new attribute. False by default.
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	overwrite   If set to true, the insert becomes a replace-insert. If a document with the same _key already exists the new document is not rejected with unique constraint violated but replaces the old document. False by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetDefaultCreateDocumentParameters() *CreateDocumentParameters {
	return &CreateDocumentParameters{
		returnNew:   false,
		returnOld:   false,
		overwrite:   false,
		waitForSync: false,
	}
}

func (req *RequestForCreateDocument) Path() string {
	return req.path
}

func (req *RequestForCreateDocument) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateDocument) Query() string {
	return ""
}

func (req *RequestForCreateDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateDocument) GetQueryParameter() string {
	return req.parameter
}

func (req *RequestForCreateDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateDocument) ResponseCode() int {
	return 202 // ok
}

//**// Response //**//

func NewResponseForCreateDocument() *ResponseForCreateDocument {
	return new(ResponseForCreateDocument)
}

type ResponseForCreateDocument []CreateResult

func (r ResponseForCreateDocument) HasResults() bool {
	if len(r) > 0 {
		return true
	} else {
		return false
	}
}

func (r ResponseForCreateDocument) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()
}

func (r *ResponseForCreateDocument) GetRawMessage() []byte {
	//return r.RawMessage
	return []byte("")
}

func (r *ResponseForCreateDocument) SetRawMessage(raw []byte) {
	//r.RawMessage = raw
}

type CreateResult struct {
	Id  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
	Rev string `json:"_rev,omitempty"`
}

func (r CreateResult) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v",
		r.Id,
		r.Key,
		r.Rev,
	)
}

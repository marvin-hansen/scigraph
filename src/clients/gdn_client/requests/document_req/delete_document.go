package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteDocument(fabric, collectionName, key string, parameters *DeleteDocumentParameters) *RequestForDeleteDocument {
	return &RequestForDeleteDocument{
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName,
			key,
		),
		parameters: fmt.Sprintf("?returnOld=%v&silent=%v&waitForSync=%v",
			parameters.returnOld,
			parameters.silent,
			parameters.waitForSync,
		),
	}
}

type RequestForDeleteDocument struct {
	path       string
	parameters string
}

type DeleteDocumentParameters struct {
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	silent      bool // If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
	waitForSync bool // If set to true, returns only after data has been synced to disk. // False by default
}

// GetDefaultDeleteDocumentParameters
//  silent If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetDefaultDeleteDocumentParameters() *DeleteDocumentParameters {
	return &DeleteDocumentParameters{
		returnOld:   false,
		silent:      true,
		waitForSync: false,
	}
}

func (req *RequestForDeleteDocument) Path() string {
	return req.path
}

func (req *RequestForDeleteDocument) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteDocument) Query() string {
	return ""
}

func (req *RequestForDeleteDocument) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteDocument) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForDeleteDocument) Payload() []byte {
	return nil
}

func (req *RequestForDeleteDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteDocument() *ResponseForDeleteDocument {
	return new(ResponseForDeleteDocument)
}

type ResponseForDeleteDocument DocumentResult

func (r *ResponseForDeleteDocument) IsResponder() {}

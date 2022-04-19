package gdn_client

import (
	r "scigraph/src/clients/gdn_client/requests/document_req"
)

// CreateNewDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func (c Client) CreateNewDocument(
	fabric string,
	collectionName string,
	silent bool,
	jsonDocument []byte,
	parameters *r.CreateDocumentParameters) (response *r.ResponseForCreateDocument, err error) {

	req := r.NewRequestForCreateDocument(fabric, collectionName, silent, jsonDocument, parameters)
	response = r.NewResponseForCreateDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

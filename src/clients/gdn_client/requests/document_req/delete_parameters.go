package document_req

type DeleteDocumentParameters struct {
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	silent      bool // If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
	waitForSync bool // If set to true, returns only after data has been synced to disk. // False by default
}

// GetDefaultDeleteDocumentParameters
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//  silent If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic. True by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetDefaultDeleteDocumentParameters() *DeleteDocumentParameters {
	return &DeleteDocumentParameters{
		returnOld:   false,
		silent:      true,
		waitForSync: false,
	}
}

//GetCustomDeleteDocumentParameters
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//  silent If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetCustomDeleteDocumentParameters(returnOld, silent, waitForSync bool) *DeleteDocumentParameters {
	return &DeleteDocumentParameters{
		returnOld:   returnOld,
		silent:      silent,
		waitForSync: waitForSync,
	}
}

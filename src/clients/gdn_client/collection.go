package gdn_client

import (
	"scigraph/src/clients/gdn_client/gdn_types"
	"scigraph/src/clients/gdn_client/requests/collection_req"
)

func (c Client) GetAllCollections(fabric string) (response *collection_req.ResponseForGetAllCollections, err error) {
	req := collection_req.NewRequestForGetAllCollections(fabric)
	response = collection_req.NewResponseForGetAllCollections()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType gdn_types.CollectionType) (err error) {
	req := collection_req.NewRequestForCreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
	response := collection_req.NewResponseForCreateNewCollection()
	if err = c.request(req, response); err != nil {
		return err
	}
	return nil
}

func (c Client) GetCollectionInfo(fabric, collectionName string) (response *collection_req.ResponseForGetCollectionInfo, err error) {
	req := collection_req.NewRequestForGetCollectionInfo(fabric, collectionName)
	response = collection_req.NewResponseForGetCollectionInfo()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateCollectionProperties updates collection properties.
// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
func (c Client) UpdateCollectionProperties(fabric, collectionName string, properties *collection_req.UpdateOptions) (response *collection_req.ResponseForUpdateCollection, err error) {
	req := collection_req.NewRequestForUpdateCollection(fabric, collectionName, properties)
	response = collection_req.NewResponseForUpdateCollection()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) TruncateCollection(fabric, collectionName string) (response *collection_req.ResponseForTruncateCollection, err error) {
	req := collection_req.NewRequestForTruncateCollection(fabric, collectionName)
	response = collection_req.NewResponseForTruncateCollection()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteCollection(fabric, collectionName string, isSystem bool) (err error) {
	req := collection_req.NewRequestForDeleteCollection(fabric, collectionName, isSystem)
	response := collection_req.NewResponseForDeleteCollection()
	if err = c.request(req, response); err != nil {
		return err
	}
	// valid response: Code: 200, Error: false, ID: 159XXXXX
	return nil
}

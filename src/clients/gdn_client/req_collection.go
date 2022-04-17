package gdn_client

import (
	"scigraph/src/clients/gdn_client/gdn_types"
	"scigraph/src/clients/gdn_client/requests"
)

func (c Client) GetAllCollections(fabric string) (response *requests.ResponseForGetAllCollections, err error) {
	req := requests.NewRequestForGetAllCollections(fabric)
	response = requests.NewResponseForGetAllCollections()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType gdn_types.CollectionType) (err error) {
	req := requests.NewRequestForCreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
	response := requests.NewResponseForCreateNewCollection()
	if err = c.request(req, response); err != nil {
		return err
	}
	return nil
}

func (c Client) GetCollectionInfo(fabric, collectionName string) (response *requests.ResponseForGetCollectionInfo, err error) {
	req := requests.NewRequestForGetCollectionInfo(fabric, collectionName)
	response = requests.NewResponseForGetCollectionInfo()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) TruncateCollection(fabric, collectionName string) (response *requests.ResponseForTruncateCollection, err error) {
	req := requests.NewRequestForTruncateCollection(fabric, collectionName)
	response = requests.NewResponseForTruncateCollection()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteCollection(fabric, collectionName string, isSystem bool) (err error) {
	req := requests.NewRequestForDeleteCollection(fabric, collectionName, isSystem)
	response := requests.NewResponseForDeleteCollection()
	if err = c.request(req, response); err != nil {
		return err
	}
	// valid response: Code: 200, Error: false, ID: 159XXXXX
	return nil
}

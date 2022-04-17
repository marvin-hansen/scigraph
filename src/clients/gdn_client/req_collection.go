package gdn_client

import "scigraph/src/clients/gdn_client/requests"

func (c Client) GetAllCollections(fabric string) (response *requests.ResponseForGetAllCollections, err error) {
	req := requests.NewRequestForGetAllCollections(fabric)
	response = requests.NewResponseForGetAllCollections()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

package document

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"scigraph/src/clients/gdn_client/requests/document_req"
	"testing"
)

const verbose = true

func getTestData() []byte {
	return []byte(` 
		[
		  {
			"item1": "data1"
		  },
		  {
			"item2": "data2"
		  }
		]
	`)
}

func TestCreateNewDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	silent := false
	jsonDocument := getTestData()

	res, err := c.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	assert.NoError(t, err)

	if verbose {
		if res != nil {
			assert.NotNil(t, res)
			for _, v := range *res {
				println(v.String())
			}
		}
	}
}

func TestDeleteDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "1"

	res, err := c.DeleteDocument(fabric, collName, key, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestDeleteDocumentNONSilent(t *testing.T) {

	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "2"
	para := document_req.GetCustomDeleteDocumentParameters(false, false, false)

	res, err := c.DeleteDocument(fabric, collName, key, para)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}

}

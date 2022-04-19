package document

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"scigraph/src/clients/gdn_client/requests/document_req"
	"testing"
)

const verbose = true

func TestCreateNewDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	silent := false
	jsonDocument := getTestInsertData()

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

	res, err = c.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	assert.NoError(t, err)

}

func TestUpdateDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	key := "2"
	jsonDocument := getTestUpdateData(key)

	res, err := c.UpdateDocument(fabric, collName, jsonDocument, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

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
}

func TestDeleteDocumentNONSilent(t *testing.T) {

	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "2"
	para := document_req.GetCustomDeleteDocumentParameters(false, false, false, false)

	res, err := c.DeleteDocument(fabric, collName, key, para)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestDeleteManyDocuments(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	k1 := "3"
	k2 := "4"
	keysToDelete := getKeysToDelete(k1, k2)

	resDel, errDel := c.DeleteManyDocuments(fabric, collName, keysToDelete, nil)
	assert.NoError(t, errDel)
	assert.NotNil(t, resDel)
}

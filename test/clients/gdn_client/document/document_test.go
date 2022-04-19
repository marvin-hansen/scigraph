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

func TestGetDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "5"

	res, err := c.GetDocument(fabric, collName, key)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printJsonRes(res, verbose)
}

func TestUpdateDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "7"
	jsonDocument := getTestUpdateSingleData(key)
	var silent = true

	res, err := c.UpdateDocument(fabric, collName, key, jsonDocument, silent, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	//printRes(res, verbose) Nothing to print as we do silent update

	silent = false
	res, err = c.UpdateDocument(fabric, collName, key, jsonDocument, silent, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)

}

func TestReplaceDocuments(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "5"

	jsonDocument := getTestReplaceSingleData(key)

	res, err := c.ReplaceDocument(fabric, collName, key, jsonDocument, nil)
	assert.NoError(t, err)
	printRes(res, verbose)

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

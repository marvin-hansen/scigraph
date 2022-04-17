package collection

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"scigraph/src/clients/gdn_client/gdn_types"
	"scigraph/src/clients/gdn_client/requests/collection_req"
	"testing"
)

const verbose = false

func TestGetAllCollections(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"

	res, err := c.GetAllCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestCreateCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collType := gdn_types.DocumentCollectionType
	collName := "TestCollection"

	err := c.CreateNewCollection(fabric, collName, false, collType)
	assert.NoError(t, err)
}

func TestGetCollectionInfo(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	res, err := c.GetCollectionInfo(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestUpdateCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	properties := &collection_req.UpdateOptions{
		// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
		HasStream:   true,
		WaitForSync: true,
	}

	res, err := c.UpdateCollectionProperties(fabric, collName, properties)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestTruncateCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	res, err := c.TruncateCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestDeleteCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	err := c.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}

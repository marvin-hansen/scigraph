package collection

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"scigraph/src/clients/gdn_client/gdn_types"
	"testing"
)

func TestGetAllCollections(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"

	res, err := c.GetAllCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	println(res.String())
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
	println(res.String())
}

func TestTruncateCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	res, err := c.TruncateCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	println(res.String())
}

func TestDeleteCollection(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	err := c.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}

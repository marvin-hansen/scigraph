package document

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"testing"
)

func TestReplaceManyDocuments(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	key := "2"
	jsonDocument := getTestReplaceData(key)

	res, err := c.ReplaceManyDocuments(fabric, collName, jsonDocument, nil)
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

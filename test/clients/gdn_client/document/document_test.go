package document

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"testing"
)

const verbose = true

func TestCreateNewDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	silent := false
	jsonDocument := []byte(` 
		[
		  {
			"item1": "data1"
		  },
		  {
			"item2": "data2"
		  }
		]
	`)

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

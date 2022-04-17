package collection

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"testing"
)

func TestGetAllCollections(t *testing.T) {
	c := gdn_client.NewClient(nil)

	fabric := "SouthEastAsia"
	_, err := c.GetAllCollections(fabric)

	//println("Print error")
	//println(err.Error())

	assert.NoError(t, err)
	//assert.NotNil(t, res)
	//println(res.String())
}

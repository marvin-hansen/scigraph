package collection

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
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

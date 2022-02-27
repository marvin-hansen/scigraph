package rest_test

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/infinity_client/rest"
	"testing"
)

func TestNewClient(t *testing.T) {

	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Should not be nil")

}

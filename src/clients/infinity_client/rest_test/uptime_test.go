package rest

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/infinity_client/rest"
	"testing"
)

func TestRequestForUptime(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Should not be nil")

	// req := server.NewRequestForUptime()

}

package unsplash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlChanged(T *testing.T) {
	assert := assert.New(T)

	var baseUrl = "https://example.com"
	SetupBaseUrl(baseUrl)

	assert.Equal(baseUrl, getEndpoint(base))
}

package nomics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNomics(t *testing.T) {
	t_apikey := "notvalidkey"
	nclient := NewNomics(t_apikey)

	assert.Equal(t, nclient.APIKey(), t_apikey, "they should be equal.")
}

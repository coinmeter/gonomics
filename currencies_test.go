package nomics

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrenciesTicker(t *testing.T) {
	nb := NewNomics(os.Getenv("NOMICS_API_KEY"))
	ctr := CurrenciesTickerRequest{PerPage: "100"}
	result, err := nb.CurrenciesTicker(&ctr)
	if err != nil {
		assert.Error(t, err, "An error occured")
	}

	t.Logf("%v\n\n", result)

	if len(result) < 1 {
		t.Fatalf("not enough result")
	}

	assert.Equal(t, result[0].Name, "Bitcoin", "They should be equal")
}

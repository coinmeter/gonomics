package nomics

import "testing"

func TestNewNomics(t *testing.T) {
	t_apikey := "notvalidkey"
	nclient := NewNomics(t_apikey)

	if nclient.APIKey() != t_apikey {
		t.Errorf("%s FAILED. Got %s, expected %s\n", "TestNewNomics", nclient.APIKey(), t_apikey)
	} else {
		t.Logf("%s PASSED. Got %s, expected %s\n", "TestNewNomics", nclient.APIKey(), t_apikey)
	}
}

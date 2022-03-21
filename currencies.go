package nomics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (nb *NomicsBase) CurrenciesTicker(ctr *CurrenciesTickerRequest) ([]CurrenciesTickerResponse, error) {
	result := []CurrenciesTickerResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", nb.APIBase(), CurrenciesTickerRoute), nil)
	if err != nil {
		return result, err
	}

	q := req.URL.Query()
	q.Add("key", nb.APIKey())
	q.Add("ids", ctr.IDs)
	q.Add("interval", ctr.Interval)
	q.Add("convert", ctr.Convert)
	q.Add("status", ctr.Status)
	q.Add("filter", ctr.Filter)
	q.Add("platform-currency", ctr.PlatformCurrency)
	q.Add("sort", ctr.Sort)
	q.Add("include-transparency", ctr.IncludeTransparncy)
	q.Add("per-page", ctr.PerPage)
	q.Add("page", ctr.Page)

	req.URL.RawQuery = q.Encode()

	resp, err := nb.Client().Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}
	return result, nil
}

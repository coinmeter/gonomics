package nomics

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MakeRequest(c *http.Client, hr *http.Request, result interface{}) (interface{}, error) {
	resp, err := c.Do(hr)
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

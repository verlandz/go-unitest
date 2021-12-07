package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	_http "net/http"
)

// assuming this is config.
var url = "http://www.randomnumberapi.com/api/v1.0/random?min=1&max=1000&count=1"

// GetRandNumber returns random number from http://www.randomnumberapi.com/.
func (r *repository) GetRandNumber(n int) (int, error) {
	if r.client == nil {
		return 0, errors.New("client can't be nil")
	}

	// http request
	req, err := _http.NewRequest(_http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	// http response
	resp, err := r.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	// data
	var data []int
	err = json.Unmarshal(b, &data)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, errors.New("empty data")
	}

	res := data[0] + n // random logic
	return res, nil
}

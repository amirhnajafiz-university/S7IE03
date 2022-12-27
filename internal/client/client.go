package client

import (
	"net/http"

	"github.com/ceit-aut/policeman/internal/model"
)

// MakeHTTPRequest on endpoint address.
func MakeHTTPRequest(endpoint model.Endpoint) (*model.Request, error) {
	// create a new http request
	req := &model.Request{
		Url: endpoint.Url,
	}

	// make http request
	resp, err := http.Get(endpoint.Url)
	if err != nil {
		return nil, err
	}

	req.Code = resp.StatusCode

	return req, nil
}

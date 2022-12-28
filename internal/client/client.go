package client

import (
	"net/http"
	"time"

	"github.com/ceit-aut/policeman/internal/model"
)

// MakeHTTPRequest on endpoint address.
func MakeHTTPRequest(endpoint model.Endpoint) (*model.Request, error) {
	// create a new http request
	req := &model.Request{
		EndpointId: endpoint.ID.String(),
		CreateTime: time.Now(),
	}

	// make http request
	resp, err := http.Get(endpoint.Url)
	if err != nil {
		return nil, err
	}

	req.Code = resp.StatusCode

	return req, nil
}

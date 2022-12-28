package response

import (
	"time"
)

type EndpointResponse struct {
	Address   string            `json:"address"`
	CreatedAt time.Time         `json:"created_at"`
	Requests  []EndpointRequest `json:"requests"`
}

type EndpointRequest struct {
	Status int       `json:"status"`
	Time   time.Time `json:"time"`
}

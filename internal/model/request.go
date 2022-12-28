package model

import (
	"time"
)

type Request struct {
	EndpointId string    `bson:"endpoint_id"`
	Code       int       `bson:"result"`
	CreateTime time.Time `json:"create_time"`
}

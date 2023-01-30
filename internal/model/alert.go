package model

import "time"

type Alert struct {
	CreateTime time.Time `bson:"create_time"`
	EndpointId string    `bson:"endpoint_id"`
	Errors     int       `json:"errors"`
}

package model

import (
	"time"
)

type Request struct {
	Url        string    `bson:"url"`
	Code       int       `bson:"result"`
	CreateTime time.Time `json:"create_time"`
}

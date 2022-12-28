package model

import (
	"time"
)

type Endpoint struct {
	Username    string    `bson:"username"`
	Url         string    `bson:"url"`
	Threshold   int       `bson:"threshold"`
	FailedTimes int       `bson:"failedTimes"`
	CreateTime  time.Time `json:"create_time"`
}

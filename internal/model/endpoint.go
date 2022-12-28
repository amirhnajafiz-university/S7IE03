package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Endpoint struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username    string             `bson:"username"`
	Url         string             `bson:"url"`
	Threshold   int                `bson:"threshold"`
	FailedTimes int                `bson:"failedTimes"`
	CreateTime  time.Time          `json:"create_time"`
}

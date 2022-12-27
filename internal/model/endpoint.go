package model

type Endpoint struct {
	Username    string `bson:"username"`
	Url         string `bson:"url"`
	Threshold   int    `bson:"threshold"`
	FailedTimes int    `bson:"failedTimes"`
}

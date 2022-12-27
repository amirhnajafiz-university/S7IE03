package model

type Request struct {
	Url  string `bson:"url"`
	Code int    `bson:"result"`
}

package worker

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Worker struct {
	db      *mongo.Database
	timeout int
}

// Start worker for monitoring endpoints.
func (w *Worker) Start() error {
	for {
		// get all endpoints
		// make http request for all
		// check the errors
		// save the request into mongodb

		time.Sleep(time.Duration(w.timeout) * time.Minute)
	}
}

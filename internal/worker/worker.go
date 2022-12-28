package worker

import (
	"time"

	"github.com/ceit-aut/policeman/internal/model"
	"github.com/ceit-aut/policeman/internal/repositories"
)

// Worker handles to check endpoints to monitor them
// and updates it status.
type Worker struct {
	repositories repositories.Repositories
	timeout      int
	workers      int
}

// New builds a new repositories' struct.
func New(r repositories.Repositories, timeout int, workers int) *Worker {
	return &Worker{
		repositories: r,
		timeout:      timeout,
		workers:      workers,
	}
}

// Start worker for monitoring endpoints.
func (w *Worker) Start() error {
	// create a channel
	channel := make(chan model.Endpoint)

	// create a new worker pool
	newPool(w.workers, channel, w.repositories)

	for {
		// get all endpoints
		endpoints := w.repositories.Endpoints.GetAll()

		// start worker for each endpoint
		for _, endpoint := range endpoints {
			channel <- endpoint
		}

		time.Sleep(time.Duration(w.timeout) * time.Minute)
	}
}

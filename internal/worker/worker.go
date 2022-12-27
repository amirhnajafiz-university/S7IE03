package worker

import (
	"time"

	"github.com/ceit-aut/policeman/internal/repositories"
)

// Worker handles to check endpoints to monitor them
// and updates it status.
type Worker struct {
	repositories repositories.Repositories
	timeout      int
}

// New builds a new repositories' struct.
func New(r repositories.Repositories, timeout int) *Worker {
	return &Worker{
		repositories: r,
		timeout:      timeout,
	}
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

package worker

import (
	"time"

	"github.com/ceit-aut/S7IE03/internal/model"
	"github.com/ceit-aut/S7IE03/internal/repositories"
)

// Worker handles to check endpoints to monitor them
// and updates it status.
type Worker struct {
	repositories repositories.Repositories
	cfg          Config
}

// New builds a new repositories' struct.
func New(cfg Config, r repositories.Repositories) *Worker {
	return &Worker{
		repositories: r,
		cfg:          cfg,
	}
}

// Start worker for monitoring endpoints.
func (w *Worker) Start() error {
	// create a channel
	channel := make(chan model.Endpoint)

	// create a new worker pool
	newPool(w.cfg.Workers, channel, w.repositories)

	for {
		// get all endpoints
		endpoints := w.repositories.Endpoints.GetAll()

		// start worker for each endpoint
		for _, endpoint := range endpoints {
			channel <- endpoint
		}

		time.Sleep(time.Duration(w.cfg.Timeout) * time.Second)
	}
}

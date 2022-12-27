package worker

import (
	"log"
	"sync"
	"time"

	"github.com/ceit-aut/policeman/internal/client"
	"github.com/ceit-aut/policeman/internal/model"
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
		// create wait group
		wg := sync.WaitGroup{}

		// get all endpoints
		endpoints := w.repositories.Endpoints.GetAll()

		// make http request for all
		for _, endpoint := range endpoints {
			wg.Add(1)

			go func(ep model.Endpoint) {
				w.do(ep)

				wg.Done()
			}(endpoint)
		}

		// wait for all requests
		wg.Wait()

		log.Println("all endpoints are update now")

		time.Sleep(time.Duration(w.timeout) * time.Minute)
	}
}

// do will check a http endpoint.
func (w *Worker) do(endpoint model.Endpoint) {
	if req, err := client.MakeHTTPRequest(endpoint); err != nil {
		log.Printf("make http request for endpoint failed:\n\t%v\n", err)
	} else {
		if e := w.repositories.Requests.Insert(*req); e != nil {
			log.Printf("failed to store request:\n\t%v\n", e)
		}
	}
}

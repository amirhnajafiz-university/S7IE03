package worker

import (
	"log"

	"github.com/ceit-aut/S7IE03/internal/client"
	"github.com/ceit-aut/S7IE03/internal/model"
	"github.com/ceit-aut/S7IE03/internal/repositories"
)

// innerWorker manages to do the tasks.
type innerWorker struct {
	repositories repositories.Repositories
	channel      chan model.Endpoint
}

// listen for events and send them to workers.
func (i *innerWorker) listen() {
	for {
		// listen on channel
		ep := <-i.channel

		// make http request
		i.do(ep)
	}
}

// do will check a http endpoint.
func (i *innerWorker) do(endpoint model.Endpoint) {
	if req, err := client.MakeHTTPRequest(endpoint); err != nil {
		log.Printf("make http request for endpoint failed:\n\t%v\n", err)
	} else {
		if e := i.repositories.Requests.Insert(*req); e != nil {
			log.Printf("failed to store request:\n\t%v\n", e)
		} else {
			if !(req.Code >= 200 && req.Code < 300) {
				endpoint.FailedTimes++

				if er := i.repositories.Endpoints.Update(endpoint); er != nil {
					log.Printf("failed to update request: \n\t%v\n", e)
				}
			}
		}
	}
}
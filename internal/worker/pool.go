package worker

import (
	"github.com/ceit-aut/S7IE03/internal/model"
	"github.com/ceit-aut/S7IE03/internal/repositories"
)

// newPool generates a new worker pool.
func newPool(capacity int, channel chan model.Endpoint, r repositories.Repositories) {
	for i := 0; i < capacity; i++ {
		iw := innerWorker{
			channel:      channel,
			repositories: r,
		}

		go iw.listen()
	}
}

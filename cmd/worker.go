package cmd

import (
	"log"

	"github.com/ceit-aut/policeman/internal/repositories"
	"github.com/ceit-aut/policeman/internal/worker"
)

type Worker struct {
	Cfg          worker.Config
	Repositories repositories.Repositories
}

func (w *Worker) Command() {

}

func (w *Worker) main() {
	// create worker
	wk := worker.New(w.Cfg, w.Repositories)

	log.Println("worker started ...")

	// start worker
	if err := wk.Start(); err != nil {
		panic(err)
	}
}

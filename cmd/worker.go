package cmd

import (
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

}

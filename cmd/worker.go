package cmd

import (
	"log"

	"github.com/ceit-aut/S7IE03/internal/repositories"
	"github.com/ceit-aut/S7IE03/internal/worker"

	"github.com/spf13/cobra"
)

type Worker struct {
	Cfg          worker.Config
	Repositories repositories.Repositories
}

func (w Worker) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { w.main() }
	return &cobra.Command{Use: "worker", Short: "start worker to monitor endpoints", Run: run}
}

func (w Worker) main() {
	// create worker
	wk := worker.New(w.Cfg, w.Repositories)

	log.Println("worker started ...")

	// start worker
	if err := wk.Start(); err != nil {
		panic(err)
	}
}

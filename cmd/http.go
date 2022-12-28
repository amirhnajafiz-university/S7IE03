package cmd

import (
	"github.com/ceit-aut/policeman/internal/config"
	"github.com/ceit-aut/policeman/internal/repositories"
)

type HTTP struct {
	Cfg          config.Config
	Repositories repositories.Repositories
}

func (h *HTTP) Command() {

}

func main() {

}
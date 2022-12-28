package handler

import (
	"github.com/ceit-aut/policeman/internal/repositories"
)

type Handler struct {
	Repositories repositories.Repositories
}

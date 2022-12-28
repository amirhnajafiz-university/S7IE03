package handler

import (
	"github.com/ceit-aut/policeman/internal/repositories"
	"github.com/ceit-aut/policeman/pkg/auth"
)

type Handler struct {
	JWT          auth.Auth
	Repositories repositories.Repositories
	Threshold    int
}

package middleware

import (
	"github.com/ceit-aut/S7IE03/internal/repositories"
	"github.com/ceit-aut/S7IE03/pkg/auth"
)

type Middleware struct {
	Repositories repositories.Repositories
	Auth         *auth.Auth
}

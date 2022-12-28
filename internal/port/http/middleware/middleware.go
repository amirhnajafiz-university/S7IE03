package middleware

import (
	"github.com/ceit-aut/policeman/internal/repositories"
	"github.com/ceit-aut/policeman/pkg/auth"
)

type Middleware struct {
	Repositories repositories.Repositories
	Auth         *auth.Auth
}

package middleware

import "github.com/ceit-aut/policeman/pkg/auth"

type Middleware struct {
	Auth *auth.Auth
}

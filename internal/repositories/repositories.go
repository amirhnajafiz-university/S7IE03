package repositories

import (
	"github.com/ceit-aut/policeman/internal/repositories/endpoints"
	"github.com/ceit-aut/policeman/internal/repositories/requests"
	"github.com/ceit-aut/policeman/internal/repositories/users"
)

// Repositories manages to keep the models repositories.
type Repositories struct {
	Users     users.Repository
	Endpoints endpoints.Repository
	Requests  requests.Repository
}

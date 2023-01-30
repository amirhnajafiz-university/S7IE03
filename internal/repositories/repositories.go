package repositories

import (
	"github.com/ceit-aut/S7IE03/internal/repositories/endpoints"
	"github.com/ceit-aut/S7IE03/internal/repositories/requests"
	"github.com/ceit-aut/S7IE03/internal/repositories/users"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repositories manages to keep the models repositories.
type Repositories struct {
	Users     users.Repository
	Endpoints endpoints.Repository
	Requests  requests.Repository
}

// New returns a repositories' struct.
func New(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:     users.New(db),
		Endpoints: endpoints.New(db),
		Requests:  requests.New(db),
	}
}

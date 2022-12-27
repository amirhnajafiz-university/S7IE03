package endpoints

import (
	"github.com/ceit-aut/policeman/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository manages the endpoints models.
type Repository interface {
	GetAll() []model.Endpoint
	GetSingle(username string) *model.Endpoint
	Upsert(endpoint model.Endpoint) error
}

type repository struct {
	db *mongo.Database
}

// New generates a new repository interface.
func New(db *mongo.Database) Repository {
	return &repository{
		db: db,
	}
}

// GetAll endpoints.
func (r *repository) GetAll() []model.Endpoint {
	return nil
}

// GetSingle endpoint by username as primary key.
func (r *repository) GetSingle(username string) *model.Endpoint {
	return nil
}

// Upsert update or insert and endpoint.
func (r *repository) Upsert(endpoint model.Endpoint) error {
	return nil
}

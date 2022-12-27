package requests

import (
	"github.com/ceit-aut/policeman/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository manages the requests models.
type Repository interface {
	GetAll(url string) []model.Request
	Insert(request model.Request) error
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

// GetAll requests.
func (r *repository) GetAll(url string) []model.Request {
	return nil
}

// Insert a new request.
func (r *repository) Insert(request model.Request) error {
	return nil
}

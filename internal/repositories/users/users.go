package users

import (
	"github.com/ceit-aut/policeman/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository manages the users models.
type Repository interface {
	GetSingle(username string) *model.User
	Insert(user model.User) error
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

// GetSingle user by username.
func (r *repository) GetSingle(username string) *model.User {
	return nil
}

// Insert a new user.
func (r *repository) Insert(user model.User) error {
	return nil
}

package users

import (
	"context"

	"github.com/ceit-aut/S7IE03/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "users"

// Repository manages the users models.
type Repository interface {
	Exists(username string) bool
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

// Exists check to see if user exists before.
func (r *repository) Exists(username string) bool {
	return r.GetSingle(username) != nil
}

// GetSingle user by username.
func (r *repository) GetSingle(username string) *model.User {
	var (
		user model.User

		ctx    = context.Background()
		filter = bson.M{"username": username}

		collection = r.db.Collection(collectionName)
	)

	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil
	}

	return &user
}

// Insert a new user.
func (r *repository) Insert(user model.User) error {
	var (
		ctx        = context.Background()
		collection = r.db.Collection(collectionName)
	)

	_, err := collection.InsertOne(ctx, user)

	return err
}

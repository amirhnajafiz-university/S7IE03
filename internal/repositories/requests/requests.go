package requests

import (
	"context"

	"github.com/ceit-aut/S7IE03/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "requests"

// Repository manages the requests models.
type Repository interface {
	GetAll(id string) []model.Request
	Insert(request model.Request) error
	Delete(id string) error
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

// GetAll requests of an endpoint.
func (r *repository) GetAll(id string) []model.Request {
	var (
		requests []model.Request
		request  model.Request

		ctx    = context.Background()
		filter = bson.M{"endpoint_id": id}

		collection = r.db.Collection(collectionName)
	)

	if cursor, err := collection.Find(ctx, filter); err == nil {
		for cursor.Next(ctx) {
			if er := cursor.Decode(&request); er == nil {
				requests = append(requests, request)
			}
		}
	}

	return requests
}

// Insert a new request.
func (r *repository) Insert(request model.Request) error {
	var (
		ctx        = context.Background()
		collection = r.db.Collection(collectionName)
	)

	_, err := collection.InsertOne(ctx, request)

	return err
}

// Delete all requests.
func (r *repository) Delete(id string) error {
	var (
		ctx    = context.Background()
		filter = bson.M{"endpoint_id": id}

		collection = r.db.Collection(collectionName)
	)

	_, err := collection.DeleteMany(ctx, filter)

	return err
}

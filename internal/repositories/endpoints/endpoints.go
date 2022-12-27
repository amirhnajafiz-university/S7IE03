package endpoints

import (
	"context"

	"github.com/ceit-aut/policeman/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "endpoints"

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
	var (
		endpoints []model.Endpoint
		endpoint  model.Endpoint

		ctx    = context.Background()
		filter = bson.D{}

		collection = r.db.Collection(collectionName)
	)

	if cursor, err := collection.Find(ctx, filter); err == nil {
		for cursor.Next(ctx) {
			if er := cursor.Decode(&endpoint); er == nil {
				endpoints = append(endpoints, endpoint)
			}
		}
	}

	return endpoints
}

// GetSingle endpoint by username as primary key.
func (r *repository) GetSingle(username string) *model.Endpoint {
	var (
		endpoint model.Endpoint

		ctx    = context.Background()
		filter = bson.M{"username": username}

		collection = r.db.Collection(collectionName)
	)

	if err := collection.FindOne(ctx, filter).Decode(&endpoint); err != nil {
		return nil
	}

	return &endpoint
}

// Upsert update or insert and endpoint.
func (r *repository) Upsert(endpoint model.Endpoint) error {
	var (
		ctx    = context.Background()
		filter = bson.M{"username": endpoint.Username}

		collection = r.db.Collection(collectionName)
	)

	_, err := collection.UpdateOne(ctx, filter, endpoint)

	return err
}

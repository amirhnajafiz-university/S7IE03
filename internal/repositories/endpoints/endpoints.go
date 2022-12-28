package endpoints

import (
	"context"
	"github.com/ceit-aut/policeman/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "endpoints"

// Repository manages the endpoints models.
type Repository interface {
	GetAll() []model.Endpoint
	GetUserEndpoints(string) []model.Endpoint
	GetSingle(string) *model.Endpoint
	Insert(model.Endpoint) (string, error)
	Update(model.Endpoint) error
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

// GetUserEndpoints person endpoints by username as primary key.
func (r *repository) GetUserEndpoints(username string) []model.Endpoint {
	var (
		endpoints []model.Endpoint
		endpoint  model.Endpoint

		ctx    = context.Background()
		filter = bson.M{"username": username}

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

// GetSingle returns one endpoint.
func (r *repository) GetSingle(id string) *model.Endpoint {
	var (
		endpoint model.Endpoint

		ctx    = context.Background()
		filter = bson.M{"_id": id}

		collection = r.db.Collection(collectionName)
	)

	if err := collection.FindOne(ctx, filter).Decode(&endpoint); err != nil {
		return nil
	}

	return &endpoint
}

// Insert an endpoint.
func (r *repository) Insert(endpoint model.Endpoint) (string, error) {
	var (
		ctx = context.Background()

		collection = r.db.Collection(collectionName)
	)

	res, err := collection.InsertOne(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

// Update on endpoint.
func (r *repository) Update(endpoint model.Endpoint) error {
	var (
		ctx = context.Background()

		filter = bson.M{"_id": endpoint.ID}
		update = bson.D{{"$set", bson.D{{"failedTimes", endpoint.FailedTimes}}}}

		collection = r.db.Collection(collectionName)
	)

	_, err := collection.UpdateOne(ctx, filter, update)

	return err
}

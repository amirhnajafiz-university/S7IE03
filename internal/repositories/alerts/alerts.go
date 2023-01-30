package alerts

import (
	"context"
	"github.com/ceit-aut/S7IE03/internal/model"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "alerts"

// Repository manages the alerts models
type Repository interface {
	GetAll(id string) []model.Alert
	Insert(alert model.Alert) error
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

// GetAll alerts of an endpoint.
func (r *repository) GetAll(id string) []model.Alert {
	var (
		alerts []model.Alert
		alert  model.Alert

		ctx    = context.Background()
		filter = bson.M{"endpoint_id": id}

		collection = r.db.Collection(collectionName)
	)

	if cursor, err := collection.Find(ctx, filter); err != nil {
		for cursor.Next(ctx) {
			if er := cursor.Decode(&alert); er == nil {
				alerts = append(alerts, alert)
			}
		}
	}

	return alerts
}

func (r *repository) Insert(alert model.Alert) error {

}
package alerts

import (
	"github.com/ceit-aut/S7IE03/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "alerts"

// Repository manages the alerts models
type Repository interface {
	GetAllForEndpoint(id string) []model.Alert
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

func (r *repository) GetAllForEndpoint(id string) []model.Alert {

}

func (r *repository) Insert(alert model.Alert) error {

}
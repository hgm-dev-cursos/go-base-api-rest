package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Example
// swagger:model
type Example struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

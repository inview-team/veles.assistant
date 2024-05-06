package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Action struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Type string             `bson:"type"`
	// Add more fields as necessary.
}

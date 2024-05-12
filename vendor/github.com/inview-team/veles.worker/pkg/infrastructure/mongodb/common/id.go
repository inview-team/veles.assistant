package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID struct {
	ID primitive.ObjectID `bson:"_id"`
}

type IDs []*ID

func (id *ID) ToEntity() string {
	return id.ID.Hex()
}

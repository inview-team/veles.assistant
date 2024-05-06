package storage

import (
	"context"

	"github.com/Korpenter/assistand/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActionStorage interface {
	GetActionByID(ctx context.Context, id string) (*entities.Action, error)
	SaveAction(ctx context.Context, action *entities.Action) error
}

type MongoActionStorage struct {
	collection *mongo.Collection
}

func NewMongoActionStorage(db *mongo.Database, collectionName string) *MongoActionStorage {
	return &MongoActionStorage{
		collection: db.Collection(collectionName),
	}
}

func (m *MongoActionStorage) GetActionByID(ctx context.Context, id string) (*entities.Action, error) {
	var action entities.Action
	objID, _ := primitive.ObjectIDFromHex(id)
	if err := m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&action); err != nil {
		return nil, err
	}
	return &action, nil
}

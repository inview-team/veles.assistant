package storage

import (
	"context"

	"github.com/inview-team/veles.assistant/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActionStorage interface {
	GetAction(ctx context.Context, id string) (*entities.Action, error)
	CreateAction(ctx context.Context, action *entities.Action) error
	UpdateAction(ctx context.Context, action *entities.Action) error
	DeleteAction(ctx context.Context, id string) error
}

type MongoActionStorage struct {
	collection *mongo.Collection
}

func NewMongoActionStorage(db *mongo.Database, collectionName string) *MongoActionStorage {
	return &MongoActionStorage{
		collection: db.Collection(collectionName),
	}
}

func (s *MongoActionStorage) GetAction(ctx context.Context, id string) (*entities.Action, error) {
	var action entities.Action
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&action)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (s *MongoActionStorage) CreateAction(ctx context.Context, action *entities.Action) error {
	_, err := s.collection.InsertOne(ctx, action)
	return err
}

func (s *MongoActionStorage) UpdateAction(ctx context.Context, action *entities.Action) error {
	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": action.ID}, bson.M{"$set": action})
	return err
}

func (s *MongoActionStorage) DeleteAction(ctx context.Context, id string) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

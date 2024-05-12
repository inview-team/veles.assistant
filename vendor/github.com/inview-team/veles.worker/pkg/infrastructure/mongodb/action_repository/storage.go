package action_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type actionRepository struct {
	client *mongodb.DBClient
	coll   *mongo.Collection
}

func NewActionRepository(client *mongodb.DBClient) entities.ActionRepository {
	coll := client.Collection("assistant", common.ActionCollectionName)
	return &actionRepository{
		client: client,
		coll:   coll,
	}
}

func (a actionRepository) Create(ctx context.Context, action *entities.Action) error {
	mAction, err := NewAction(action)
	if err != nil {
		return fmt.Errorf("creating action got error: %v", err)
	}

	if _, err := a.coll.InsertOne(ctx, mAction); err != nil {
		return fmt.Errorf("creating action got error: %v", err)
	}
	return nil
}

func (a actionRepository) GetByID(ctx context.Context, id string) (*entities.Action, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("finding action by id got err: %v", err)
	}

	f := bson.D{{"_id", oID}}
	res := a.coll.FindOne(ctx, f)

	var mAction Action
	if err = res.Decode(&mAction); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entities.ErrActionNotFound
		}
		return nil, fmt.Errorf("finding action by id got error: %v", err)
	}

	return mAction.ToEntity()
}

func (a actionRepository) NextID(_ context.Context) entities.ActionID {
	return entities.ActionID(primitive.NewObjectID().Hex())
}

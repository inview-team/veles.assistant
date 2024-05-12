package scenario_repository

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

type scenarioRepository struct {
	client *mongodb.DBClient
	coll   *mongo.Collection
}

func NewScenarioRepository(client *mongodb.DBClient) entities.ScenarioRepository {
	coll := client.Collection("assistant", common.ScenarioCollectionName)
	return &scenarioRepository{
		client: client,
		coll:   coll,
	}
}

func (s scenarioRepository) Create(ctx context.Context, scenario *entities.Scenario) error {
	mScenario, err := NewScenario(scenario)
	if err != nil {
		return fmt.Errorf("creating scenario got error: %v", err)
	}

	if _, err := s.coll.InsertOne(ctx, mScenario); err != nil {
		return fmt.Errorf("creating scenario got error: %v", err)
	}
	return nil
}

func (s scenarioRepository) GetById(ctx context.Context, id string) (*entities.Scenario, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("finding scenario by id got error: %v", err)
	}

	f := bson.D{{"_id", oID}}
	res := s.coll.FindOne(ctx, f)

	var mScenario Scenario
	if err = res.Decode(&mScenario); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entities.ErrScenarioNotFound
		}
		return nil, fmt.Errorf("finding scenario by id got error: %v", err)
	}

	return mScenario.ToEntity()
}

func (s scenarioRepository) Delete(ctx context.Context, id string) (*entities.Scenario, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("deleting scenario by id got error: %v", err)
	}

	f := bson.D{{"_id", oID}}
	res := s.coll.FindOneAndDelete(ctx, f)

	var mScenario Scenario
	if err = res.Decode(&mScenario); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entities.ErrScenarioNotFound
		}
		return nil, fmt.Errorf("deleting scenario by id got error: %v", err)
	}

	return mScenario.ToEntity()
}

func (s scenarioRepository) NextID(_ context.Context) entities.ScenarioID {
	return entities.ScenarioID(primitive.NewObjectID().Hex())
}

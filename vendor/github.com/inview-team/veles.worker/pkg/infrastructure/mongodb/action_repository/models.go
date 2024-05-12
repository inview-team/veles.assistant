package action_repository

import (
	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Action struct {
	ID               primitive.ObjectID           `bson:"_id"`
	Name             string                       `bson:"name"`
	Type             int                          `bson:"type"`
	Input            map[string]entities.Variable `bson:"input"`
	Output           []string                     `bson:"output"`
	AdditionalParams map[string]interface{}       `bson:"additional_params"`
}

func (a *Action) ToEntity() (*entities.Action, error) {
	return entities.NewAction(entities.ActionID(a.ID.Hex()), a.Name, entities.ActionType(a.Type), a.Input, a.Output, a.AdditionalParams)
}

func NewAction(action *entities.Action) (*Action, error) {
	id, err := primitive.ObjectIDFromHex(string(action.Id))
	if err != nil {
		return nil, err
	}

	return &Action{
		ID:               id,
		Name:             action.Name,
		Type:             int(action.Type),
		Input:            action.Input,
		Output:           action.Output,
		AdditionalParams: action.AdditionalParams,
	}, nil
}

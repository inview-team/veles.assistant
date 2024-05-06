package service

import (
	"context"

	"github.com/Korpenter/assistand/internal/entities"
	"github.com/Korpenter/assistand/internal/storage"
	"github.com/joeycumines/go-behaviortree"
)

type MatchService interface {
	MatchAction(ctx context.Context, session *entities.Session, text string) (string, error)
}

type MatchServiceImpl struct {
	actionStorage storage.ActionStorage
}

func NewMatchService(actionStorage storage.ActionStorage) MatchService {
	return &MatchServiceImpl{actionStorage: actionStorage}
}

func (ms *MatchServiceImpl) MatchAction(ctx context.Context, session *entities.Session, text string) (string, error) {
	action, err := ms.actionStorage.GetActionByID(ctx, "some-action-id")
	if err != nil {
		return "", err
	}

	tick := behaviortree.New(func(children []behaviortree.Node) behaviortree.Status {
		return behaviortree.Success
	})

	tick.Tick()

	return "Action executed successfully", nil
}

package service

import (
	"context"
	"fmt"

	"github.com/inview-team/veles.assistant/internal/entities"
	"github.com/inview-team/veles.assistant/internal/storage"
	"github.com/inview-team/veles.assistant/pkg/proto/gen/pb"
	"google.golang.org/grpc"
)

type MatchService interface {
	ProcessMessage(ctx context.Context, session *entities.Session, text string) (*entities.Action, error)
}

type MatchServiceImpl struct {
	actionStorage storage.ActionStorage
	grpcClient    pb.MatcherClient
}

func NewMatchService(actionStorage storage.ActionStorage, grpcConn *grpc.ClientConn) MatchService {
	return &MatchServiceImpl{
		actionStorage: actionStorage,
		grpcClient:    pb.NewMatcherClient(grpcConn),
	}
}

func (ms *MatchServiceImpl) ProcessMessage(ctx context.Context, session *entities.Session, text string) (*entities.Action, error) {
	if session.State == "" {
		req := &pb.MatchScenarioRequest{
			UserPrompt: text,
		}
		res, err := ms.grpcClient.MatchScenario(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("failed to match scenario: %v", err)
		}

		if res.RootId == "" {
			return nil, fmt.Errorf("no matching scenario found")
		}
		session.State = res.RootId

		action, err := ms.actionStorage.GetAction(ctx, res.RootId)
		if err != nil {
			return nil, fmt.Errorf("failed to get action: %v", err)
		}

		return action, nil
	} else {
		action, err := ms.actionStorage.GetAction(ctx, session.State)
		if err != nil {
			return nil, fmt.Errorf("failed to get action: %v", err)
		}

		return action, nil
	}
}

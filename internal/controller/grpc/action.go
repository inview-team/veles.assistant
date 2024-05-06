package grpc

import (
	"context"
	"fmt"

	"github.com/Korpenter/assistand/internal/controller/grpc/pb"
	"github.com/Korpenter/assistand/internal/service"
)

type ActionHandler struct {
	pb.UnimplementedActionServiceServer
	matchService   service.MatchService
	sessionService service.SessionService
	executeService service.ExecuteService
}

func NewActionHandler(matcher service.MatchService, session service.SessionService, executor service.ExecuteService) *ActionHandler {
	return &ActionHandler{
		matchService:   matcher,
		sessionService: session,
		executeService: executor,
	}
}

func (h *ActionHandler) SendVoiceInput(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	session, err := h.sessionService.GetSession(req.SessionId)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %v", err)
	}

	action, err := h.matchService.MatchAction(session, req.TextInput)
	if err != nil {
		return nil, fmt.Errorf("failed to match action: %v", err)
	}

	result, err := h.executeService.ExecuteAction(session, action)
	if err != nil {
		return nil, fmt.Errorf("failed to execute action: %v", err)
	}

	return &pb.ActionResponse{
		Message: result,
		Success: true,
	}, nil
}

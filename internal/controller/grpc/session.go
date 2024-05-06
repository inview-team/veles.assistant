// session_handler.go
package grpc

import (
	"context"
	"fmt"

	"github.com/Korpenter/assistand/internal/controller/grpc/pb"
	"github.com/Korpenter/assistand/internal/service"
)

type SessionHandler struct {
	pb.UnimplementedSessionHandlerServer
	sessionService service.SessionService
}

func NewSessionHandler(sessionService service.SessionService) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) StartSession(ctx context.Context, req *pb.StartSessionRequest) (*pb.StartSessionResponse, error) {
	sessionID, err := h.sessionService.StartSession(req.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	return &pb.StartSessionResponse{
		SessionId: sessionID,
	}, nil
}

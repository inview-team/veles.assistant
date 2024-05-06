// session_handler.go
package grpc

import (
	"context"
	"fmt"

	"github.com/Korpenter/assistand/internal/controller/grpc/pb"
	"github.com/Korpenter/assistand/internal/service"
)

type SessionHandler struct {
	pb.UnimplementedSessionServiceServer
	sessionService service.SessionService
}

func NewSessionHandler(sessionService service.SessionService) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) CreateSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	sessionID, err := h.sessionService.CreateSession(req.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	return &pb.SessionResponse{
		SessionId: sessionID,
		Success:   true,
	}, nil
}

func (h *SessionHandler) GetSession(ctx context.Context, req *pb.SessionRequest) (*pb.SessionResponse, error) {
	session, err := h.sessionService.GetSession(req.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session: %v", err)
	}
	return &pb.SessionResponse{
		SessionId: session.ID,
		State:     session.State,
		Success:   true,
	}, nil
}

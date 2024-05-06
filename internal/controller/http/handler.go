package http

import (
	"github.com/Korpenter/assistand/internal/service"
	"github.com/gorilla/websocket"
)

type HttpHandler struct {
	sessionService service.SessionService
	matchService   service.MatchService
	executeService service.ExecuteService
	Upgrader       websocket.Upgrader
}

func NewHttpHandler(ss service.SessionService, ms service.MatchService, es service.ExecuteService) *HttpHandler {
	handler := &HttpHandler{
		sessionService: ss,
		executeService: es,
		matchService:   ms,
	}
	return handler
}

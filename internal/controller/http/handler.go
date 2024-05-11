package http

import (
	"github.com/gorilla/websocket"
	"github.com/inview-team/veles.assistant/internal/service"
)

type HttpHandler struct {
	sessionService service.SessionService
	matchService   service.MatchService
	executeService service.ExecuteService
	Upgrader       websocket.Upgrader
}

func NewHttpHandler(ss service.SessionService, ms service.MatchService, es service.ExecuteService) *HttpHandler {
	return &HttpHandler{
		sessionService: ss,
		matchService:   ms,
		executeService: es,
	}
}

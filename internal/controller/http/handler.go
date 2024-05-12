package http

import (
	"encoding/json"
	"net/http"

	"github.com/inview-team/veles.assistant/internal/hub"
	"github.com/inview-team/veles.assistant/internal/service"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

type HttpHandler struct {
	sessionService service.SessionService
	hub            hub.Hub
	actionService  service.ActionService
}

func NewHttpHandler(ss service.SessionService, as service.ActionService, hub hub.Hub) *HttpHandler {
	return &HttpHandler{
		hub:            hub,
		sessionService: ss,
		actionService:  as,
	}
}

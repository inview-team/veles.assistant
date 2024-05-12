package http

import (
	"encoding/json"
	"net/http"

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
	matchService   service.MatchService
	executeService service.ExecuteService
}

func NewHttpHandler(ss service.SessionService, ms service.MatchService, es service.ExecuteService) *HttpHandler {
	return &HttpHandler{
		sessionService: ss,
		matchService:   ms,
		executeService: es,
	}
}

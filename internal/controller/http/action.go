package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/inview-team/veles.assistant/pkg/common"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) HandleAction(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		jsonResponse(w, http.StatusBadRequest, common.ErrorResponse{Error: "missing Authorization header"})
		return
	}

	token := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		jsonResponse(w, http.StatusBadRequest, common.ErrorResponse{Error: "invalid Authorization header format"})
		return
	}

	var req common.ActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, common.ErrorResponse{Error: fmt.Sprintf("invalid request: %v", err)})
		return
	}

	session, err := h.sessionService.GetSession(req.SessionID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to get session: %v", err)})
		return
	}

	session.Token = token
	go func() {
		output, scenarioID, jobID, err := h.actionService.ProcessMessage(context.Background(), session, req.Action)
		if err != nil {
			log.Errorf("failed to process message: %v", err)
		}

		err = h.sessionService.UpdateSessionState(session.ID, scenarioID, jobID)
		if err != nil {
			log.Errorf("failed to update session state: %v", err)
		}
		outHub, err := json.Marshal(common.ActionResponse{State: jobID, Text: output})
		log.Info("sending to hub err: ", err)
		log.Info("sending to hub session: ", output, jobID, scenarioID, session.ID)
		h.hub.SendMessage(session.ID, outHub)
	}()
	jsonResponse(w, http.StatusOK, common.ActionResponse{State: "", Text: "received"})
}

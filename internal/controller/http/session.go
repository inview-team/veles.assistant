package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inview-team/veles.assistant/pkg/common"
)

func (h *HttpHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	var req common.InitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, common.ErrorResponse{Error: fmt.Sprintf("invalid request: %v", err)})
		return
	}

	if req.SessionID == "" {
		sessionID, err := h.sessionService.StartSession(req.Token)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("error creating session: %v", err)})
			return
		}
		jsonResponse(w, http.StatusOK, common.InitResponse{SessionID: sessionID})
		return
	}

	session, err := h.sessionService.GetSession(req.SessionID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("error retrieving session: %v", err)})
		return
	}
	jsonResponse(w, http.StatusOK, common.InitResponse{SessionID: session.ID, State: session.State})
}

func (h *HttpHandler) UpdateSessionToken(w http.ResponseWriter, r *http.Request) {
	var req common.UpdateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, common.ErrorResponse{Error: fmt.Sprintf("invalid request: %v", err)})
		return
	}

	if err := h.sessionService.UpdateSessionToken(req.SessionID, req.Token); err != nil {
		jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("error updating session token: %v", err)})
		return
	}

	jsonResponse(w, http.StatusOK, nil)
}

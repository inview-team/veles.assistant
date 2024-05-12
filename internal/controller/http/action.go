package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inview-team/veles.assistant/pkg/common"
)

func (h *HttpHandler) HandleAction(w http.ResponseWriter, r *http.Request) {
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

	action, err := h.matchService.ProcessMessage(r.Context(), session, req.Action)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to process message: %v", err)})
		return
	}

	err = h.sessionService.UpdateSessionState(session.ID, action.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, common.ErrorResponse{Error: fmt.Sprintf("failed to update session state: %v", err)})
		return
	}

	jsonResponse(w, http.StatusOK, common.ActionResponse{ActionID: action.ID, State: session.State})
}

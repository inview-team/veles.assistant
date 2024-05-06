package http

import (
	"encoding/json"
	"net/http"
)

type ActionRequest struct {
	SessionID string `json:"session_id"`
	Action    string `json:"action"`
}

type ActionResult struct {
	Result string `json:"result"`
}

func (h *HttpHandler) HandleAction(w http.ResponseWriter, r *http.Request) {
	var request ActionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	session, err := h.sessionService.GetSession(request.SessionID)
	if err != nil {
		http.Error(w, "Session Not Found", http.StatusNotFound)
		return
	}

	action, err := h.matchService.MatchAction(session, request.Action)
	if err != nil {
		http.Error(w, "Action Match Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := h.executeService.ExecuteAction(session, action)
	if err != nil {
		http.Error(w, "Action Execution Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := ActionResult{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package http

import (
	"encoding/json"
	"net/http"
)

type SessionCreateRequest struct {
	Token string `json:"token"`
}

type CreateSessionResponse struct {
	SessionID string `json:"session_id"`
	State     string `json:"state"`
}

type SessionUpdateTokenRequest struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
}

func (h *HttpHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	var request SessionCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	sessionID, err := h.sessionService.StartSession(request.Token)
	if err != nil {
		http.Error(w, "Unable to Create Session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := CreateSessionResponse{
		SessionID: sessionID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *HttpHandler) UpdateSessionToken(w http.ResponseWriter, r *http.Request) {
	var request SessionUpdateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.sessionService.UpdateSessionToken(request.SessionID, request.Token)
	if err != nil {
		http.Error(w, "Unable to Update Session Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

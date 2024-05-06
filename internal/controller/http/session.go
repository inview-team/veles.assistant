package http

import "net/http"

type SessionCreateRequest struct {
	Token string `json:"token"`
}

type SessionResponse struct {
	SessionID string `json:"session_id"`
	State     string `json:"state"`
}

func (h *HttpHandler) InitSession(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *HttpHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	//
}

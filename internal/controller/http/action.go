package http

import "net/http"

type ActionRequest struct {
	SessionID string `json:"session_id"`
	Action    string `json:"action"`
}

type ActionResult struct {
	Result string `json:"result"`
}

func (h *HttpHandler) HandleAction(w http.ResponseWriter, r *http.Request) {
	//
}

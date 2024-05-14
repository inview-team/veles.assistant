package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/inview-team/veles.assistant/pkg/common"
)

func (h *WsHandler) startSession(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	var payload common.InitRequest
	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return jsonResponse(http.StatusBadRequest, fmt.Sprintf("error unmarshalling init payload: %v", err), nil)
	}

	conn := NewWebSocketConnection(wsConn)

	if payload.SessionID == "" {
		sessionID, err := h.sessionService.StartSession(payload.Token)
		if err != nil {
			return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("error creating session: %v", err), nil)
		}
		h.hub.Register(sessionID, conn)
		return jsonResponse(http.StatusOK, "session started", common.InitResponse{SessionID: sessionID, State: "Приввет! Я могу помочь узнать баланс или перевести средства!"})
	}

	session, err := h.sessionService.GetSession(payload.SessionID)
	if err != nil {
		return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("error retrieving session: %v", err), nil)
	}
	h.hub.Register(session.ID, conn)
	return jsonResponse(http.StatusOK, "session retrieved", common.InitResponse{SessionID: session.ID, State: session.JobID})
}

func (h *WsHandler) updateSessionToken(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	var payload common.UpdateTokenRequest
	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return jsonResponse(http.StatusBadRequest, fmt.Sprintf("error unmarshalling update token payload: %v", err), nil)
	}

	if err := h.sessionService.UpdateSessionToken(payload.SessionID, payload.Token); err != nil {
		return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("error updating session token: %v", err), nil)
	}

	return jsonResponse(http.StatusOK, "session token updated", nil)
}

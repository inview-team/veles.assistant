package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *WsHandler) initSession(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	var payload InitPayload
	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return nil, fmt.Errorf("error unmarshalling init payload: %v", err)
	}

	conn := NewWebSocketConnection(wsConn)

	if payload.SessionID == "" {
		sessionID, err := h.sessionService.CreateSession(payload.Token)
		if err != nil {
			return nil, fmt.Errorf("error creating session: %v", err)
		}
		h.hub.Register(sessionID, conn)
		return []byte(fmt.Sprintf("session_id: %s", sessionID)), nil
	}

	session, err := h.sessionService.GetSession(payload.SessionID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving session: %v", err)
	}
	h.hub.Register(session.ID, conn)
	return []byte(fmt.Sprintf("session_id: %s, state: %s", session.ID, session.State)), nil
}

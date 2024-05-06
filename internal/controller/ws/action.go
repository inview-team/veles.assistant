package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *WsHandler) handleAction(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	var payload ActionPayload
	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return nil, fmt.Errorf("error unmarshalling action payload: %v", err)
	}

	session, err := h.sessionService.GetSession(payload.SessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %v", err)
	}

	action, err := h.matchService.MatchAction(session, payload.Action)
	if err != nil {
		return nil, fmt.Errorf("failed to match action: %v", err)
	}

	result, err := h.executeService.ExecuteAction(session, action)
	if err != nil {
		return nil, fmt.Errorf("failed to execute action: %v", err)
	}

	return json.Marshal(result)
}

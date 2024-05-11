package ws

import (
	"context"
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

	action, err := h.matchService.ProcessMessage(context.Background(), session, payload.Action)
	if err != nil {
		return nil, fmt.Errorf("failed to process message: %v", err)
	}

	// Update session with new state (if necessary)
	err = h.sessionService.UpdateSessionState(session.ID, action.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update session: %v", err)
	}

	return json.Marshal(action)
}

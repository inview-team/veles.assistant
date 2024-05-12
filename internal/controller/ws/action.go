package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/inview-team/veles.assistant/pkg/common"
)

func (h *WsHandler) handleAction(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	var payload common.ActionRequest
	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return jsonResponse(http.StatusBadRequest, fmt.Sprintf("error unmarshalling action payload: %v", err), nil)
	}

	session, err := h.sessionService.GetSession(payload.SessionID)
	if err != nil {
		return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("failed to get session: %v", err), nil)
	}

	output, scenarioID, jobID, err := h.actionService.ProcessMessage(context.Background(), session, payload.Action)
	if err != nil {
		return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("failed to process message: %v", err), nil)
	}

	err = h.sessionService.UpdateSessionState(session.ID, scenarioID, jobID)
	if err != nil {
		return jsonResponse(http.StatusInternalServerError, fmt.Sprintf("failed to update session state: %v", err), nil)
	}

	return jsonResponse(http.StatusOK, "action processed", common.ActionResponse{State: jobID, Text: output})
}

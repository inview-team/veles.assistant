package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/inview-team/veles.assistant/internal/hub"
	"github.com/inview-team/veles.assistant/internal/service"
	"github.com/inview-team/veles.assistant/pkg/common"
	log "github.com/sirupsen/logrus"
)

const (
	wsBufferSizeLimitInBytes = 1024
)

type WsHandler struct {
	sessionService service.SessionService
	actionService  service.ActionService
	hub            hub.Hub
	handlers       map[string]func(*Request, *websocket.Conn) ([]byte, error)
	Upgrader       websocket.Upgrader
}

func NewWsHandler(ss service.SessionService, as service.ActionService, hub hub.Hub) *WsHandler {
	handler := &WsHandler{
		sessionService: ss,
		actionService:  as,
		hub:            hub,
		handlers:       make(map[string]func(*Request, *websocket.Conn) ([]byte, error)),
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  wsBufferSizeLimitInBytes,
			WriteBufferSize: wsBufferSizeLimitInBytes,
		},
	}
	handler.initHandlers()
	return handler
}

func (h *WsHandler) HandleWs(w http.ResponseWriter, req *http.Request) {
	h.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsConn, err := h.Upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to upgrade to a WS connection: %s", err.Error()))
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			log.Error(fmt.Sprintf("Unable to gracefully close WS connection: %s", err.Error()))
		}
	}(wsConn)

	log.Info("Websocket connection established")
	var mu sync.Mutex
	for {
		msgType, message, err := wsConn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Info("Closing WS connection gracefully")
			} else {
				log.Error(fmt.Sprintf("Unable to read WS message: %s", err.Error()))
				log.Info("Closing WS connection with error")
			}
			break
		}

		if msgType == websocket.TextMessage {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				resp, err := h.HandleMessage(message, wsConn)
				if err != nil {
					log.Error(fmt.Sprintf("Unable to handle WS request: %s", err.Error()))
					_ = wsConn.WriteMessage(msgType, []byte(fmt.Sprintf("WS Handle error: %s", err.Error())))
				} else {
					_ = wsConn.WriteMessage(msgType, resp)
				}
			}()
		}
	}
}

func (h *WsHandler) initHandlers() {
	h.handlers["init"] = h.startSession
	h.handlers["action"] = h.handleAction
	h.handlers["default"] = h.handleDefault
	h.handlers["update_token"] = h.updateSessionToken
}

func (h *WsHandler) HandleMessage(rawMsg []byte, wsConn *websocket.Conn) ([]byte, error) {
	var msg Request
	if err := json.Unmarshal(rawMsg, &msg); err != nil {
		return jsonResponse(http.StatusBadRequest, fmt.Sprintf("error unmarshalling message: %v", err), nil)
	}

	handlerFunc, ok := h.handlers[msg.Type]
	if !ok {
		handlerFunc = h.handleDefault
	}
	return handlerFunc(&msg, wsConn)
}

func (h *WsHandler) handleDefault(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	return jsonResponse(http.StatusBadRequest, "unhandled message type", nil)
}

func jsonResponse(status int, message string, data interface{}) ([]byte, error) {
	resp := common.NewJsonResponse(status, message, data)
	return json.Marshal(resp)
}

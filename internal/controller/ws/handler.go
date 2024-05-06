package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Korpenter/assistand/internal/hub"
	"github.com/Korpenter/assistand/internal/service"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	wsBufferSizeLimitInBytes = 1024
)

type WsHandler struct {
	sessionService service.SessionService
	matchService   service.MatchService
	executeService service.ExecuteService
	hub            hub.Hub
	handlers       map[string]func(*Request, *websocket.Conn) ([]byte, error)
	Upgrader       websocket.Upgrader
}

func NewWsHandler(ss service.SessionService, ms service.MatchService, es service.ExecuteService, hub hub.Hub) *WsHandler {
	handler := &WsHandler{
		sessionService: ss,
		executeService: es,
		matchService:   ms,
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
		log.Error(fmt.Sprintf("Unable to upgrade to a WS connection, %s", err.Error()))
		return
	}

	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			log.Error(fmt.Sprintf("Unable to gracefully close WS connection, %s", err.Error()))
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
				log.Error(fmt.Sprintf("Unable to read WS message, %s", err.Error()))
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
					log.Error(fmt.Sprintf("Unable to handle WS request, %s", err.Error()))
					_ = wsConn.WriteMessage(msgType, []byte(fmt.Sprintf("WS Handle error: %s", err.Error())))
				} else {
					_ = wsConn.WriteMessage(msgType, resp)
				}
			}()
		}
	}
}

func (h *WsHandler) initHandlers() {
	h.handlers["init"] = h.initSession
	h.handlers["action"] = h.handleAction
	h.handlers["default"] = h.handleDefault
}

func (h *WsHandler) HandleMessage(rawMsg []byte, wsConn *websocket.Conn) ([]byte, error) {
	var msg *Request
	if err := json.Unmarshal(rawMsg, &msg); err != nil {
		return nil, fmt.Errorf("error unmarshalling message: %v", err)
	}

	handlerFunc, ok := h.handlers[msg.Type]
	if !ok {
		handlerFunc = h.handlers["default"]
	}
	return handlerFunc(msg, wsConn)
}

func (h *WsHandler) handleDefault(msg *Request, wsConn *websocket.Conn) ([]byte, error) {
	return []byte("Unhandled message type"), nil
}

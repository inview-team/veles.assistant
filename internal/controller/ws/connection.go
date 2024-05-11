package ws

import (
	"github.com/gorilla/websocket"
	"github.com/inview-team/veles.assistant/internal/hub"
)

type WebSocketConnection struct {
	conn *websocket.Conn
}

func (wc *WebSocketConnection) SendMessage(data []byte) error {
	return wc.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (wc *WebSocketConnection) Close() error {
	return wc.conn.Close()
}

func NewWebSocketConnection(conn *websocket.Conn) hub.Connection {
	return &WebSocketConnection{conn: conn}
}

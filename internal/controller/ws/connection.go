package ws

import (
	"github.com/Korpenter/assistand/internal/hub"
	"github.com/gorilla/websocket"
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

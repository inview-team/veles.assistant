package hub

import (
	"fmt"
	"sync"
)

type Hub interface {
	Register(sessionID string, conn Connection)
	Unregister(sessionID string)
	SendMessage(sessionID string, message []byte) error
}

type Connection interface {
	SendMessage(message []byte) error
	Close() error
}

type MemHub struct {
	connections map[string]Connection
	mu          sync.Mutex
}

func NewMemHub() Hub {
	return &MemHub{
		connections: make(map[string]Connection),
	}
}

func (h *MemHub) Register(sessionID string, conn Connection) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.connections[sessionID] = conn
}

func (h *MemHub) Unregister(sessionID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if conn, ok := h.connections[sessionID]; ok {
		conn.Close()
		delete(h.connections, sessionID)
	}
}

func (h *MemHub) SendMessage(sessionID string, message []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	conn, ok := h.connections[sessionID]
	if !ok {
		return fmt.Errorf("no connection found for session_id %s", sessionID)
	}
	return conn.SendMessage(message)
}

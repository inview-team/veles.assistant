package ws

import "encoding/json"

type Request struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type InitPayload struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id,omitempty"`
}

type ActionPayload struct {
	SessionID string `json:"session_id"`
	Action    string `json:"action"`
}

type UpdateTokenPayload struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
}

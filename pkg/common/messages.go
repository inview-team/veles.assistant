package common

type InitRequest struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id,omitempty"`
}

type InitResponse struct {
	SessionID string `json:"session_id"`
	State     string `json:"state,omitempty"`
}

type ActionRequest struct {
	SessionID string `json:"session_id"`
	Action    string `json:"action"`
}

type ActionResponse struct {
	State string `json:"state"`
	Text  string `json:"text"`
}

type UpdateTokenRequest struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewJsonResponse(status int, message string, data interface{}) JsonResponse {
	return JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

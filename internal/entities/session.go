package entities

type Session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	State string `json:"state"`
}

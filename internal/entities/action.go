package entities

import "net/http"

type ActionNode struct {
	Name     string
	Action   func(*Session, map[string]string) (string, error)
	Children []*ActionNode
}

type Action struct {
	URL           string
	Method        string
	Headers       map[string]string
	BodyTemplate  string
	ResponseParse func(http.Response) (string, map[string]string)
}

type BackendService struct {
	URL           string
	Method        string
	Headers       map[string]string
	ResponseParse func(http.Response) (string, map[string]string)
}

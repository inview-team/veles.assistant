package entities

import (
	"context"
	"errors"
)

type Action struct {
	Id               ActionID
	Name             string
	Type             ActionType
	Input            map[string]Variable
	Output           []string
	AdditionalParams map[string]interface{}
}

type ActionID string
type ActionType int

const (
	Request ActionType = iota + 1
	Validate
)

type Variable struct {
	Value interface{}
	Type  string
}

type HTTPRequest struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    map[string]interface{}
}

type ActionRepository interface {
	Create(ctx context.Context, action *Action) error
	GetByID(ctx context.Context, id string) (*Action, error)
	NextID(ctx context.Context) ActionID
}

func NewAction(id ActionID, name string, actionType ActionType, arguments map[string]Variable, output []string, params map[string]interface{}) (*Action, error) {
	return &Action{
		Id:               id,
		Name:             name,
		Type:             actionType,
		Input:            arguments,
		Output:           output,
		AdditionalParams: params,
	}, nil
}

var (
	ErrActionNotFound = errors.New("action not found")
)

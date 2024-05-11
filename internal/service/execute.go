package service

import (
	"github.com/inview-team/veles.assistant/internal/entities"
)

type ExecuteService interface {
	ExecuteAction(session *entities.Session, action string) (string, error)
}

type ExecuteServiceImpl struct {
}

func NewExecuteService() ExecuteService {
	return &ExecuteServiceImpl{}
}

func (es *ExecuteServiceImpl) ExecuteAction(session *entities.Session, action string) (string, error) {
	return "", nil
}

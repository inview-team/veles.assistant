package service

import "github.com/inview-team/veles.assistant/internal/models"

type ExecuteService interface {
	ExecuteAction(session *models.Session, action string) (string, error)
}

type ExecuteServiceImpl struct {
}

func NewExecuteService() ExecuteService {
	return &ExecuteServiceImpl{}
}

func (es *ExecuteServiceImpl) ExecuteAction(session *models.Session, action string) (string, error) {
	return "", nil
}

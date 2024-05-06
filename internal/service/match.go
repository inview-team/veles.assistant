package service

import "github.com/Korpenter/assistand/internal/entities"

type MatchService interface {
	MatchAction(session *entities.Session, text string) (string, error)
}

type MatchServiceImpl struct {
}

func NewMatchService() MatchService {
	return &MatchServiceImpl{}
}

func (ms *MatchServiceImpl) MatchAction(session *entities.Session, text string) (string, error) {
	return "", nil
}

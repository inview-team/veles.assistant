package service

import (
	"fmt"

	"github.com/Korpenter/assistand/internal/entities"
	"github.com/Korpenter/assistand/internal/storage"
	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(token string) (string, error)
	GetSession(sessionID string) (*entities.Session, error)
	CloseSession(sessionID string) error
}

type SessionServiceImpl struct {
	storage storage.SessionStorage
}

func NewSessionService(storage storage.SessionStorage) SessionService {
	return &SessionServiceImpl{
		storage: storage,
	}
}

func (ss *SessionServiceImpl) CreateSession(token string) (string, error) {
	session := &entities.Session{
		ID:    uuid.New().String(),
		Token: token,
	}
	if err := ss.storage.CreateSession(session); err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	return session.ID, nil
}

func (ss *SessionServiceImpl) GetSession(sessionID string) (*entities.Session, error) {
	return ss.storage.GetSession(sessionID)
}

func (ss *SessionServiceImpl) CloseSession(sessionID string) error {
	return ss.storage.DeleteSession(sessionID)
}

package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/inview-team/veles.assistant/internal/models"
	"github.com/inview-team/veles.assistant/internal/storage"
)

type SessionService interface {
	StartSession(token string) (string, error)
	UpdateSessionToken(sessionID, token string) error
	UpdateSessionState(sessionID, scenarioID, jobID string) error
	GetSession(sessionID string) (*models.Session, error)
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

func (ss *SessionServiceImpl) StartSession(token string) (string, error) {
	session := &models.Session{
		ID:    uuid.New().String(),
		Token: token,
	}
	if err := ss.storage.CreateSession(session); err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	return session.ID, nil
}

func (ss *SessionServiceImpl) UpdateSessionToken(sessionID, token string) error {
	session, err := ss.storage.GetSession(sessionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve session: %v", err)
	}
	session.Token = token
	if err := ss.storage.UpdateSession(session); err != nil {
		return fmt.Errorf("failed to update session: %v", err)
	}
	return nil
}

func (ss *SessionServiceImpl) UpdateSessionState(sessionID, scenarioID, jobID string) error {
	session, err := ss.storage.GetSession(sessionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve session: %v", err)
	}
	session.ScenarioID = scenarioID
	session.JobID = jobID
	if err := ss.storage.UpdateSession(session); err != nil {
		return fmt.Errorf("failed to update session: %v", err)
	}
	return nil
}

func (ss *SessionServiceImpl) GetSession(sessionID string) (*models.Session, error) {
	return ss.storage.GetSession(sessionID)
}

func (ss *SessionServiceImpl) CloseSession(sessionID string) error {
	return ss.storage.DeleteSession(sessionID)
}

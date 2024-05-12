package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/inview-team/veles.assistant/internal/models"
	"github.com/redis/go-redis/v9"
)

type SessionStorage interface {
	CreateSession(session *models.Session) error
	UpdateSession(session *models.Session) error
	GetSession(sessionID string) (*models.Session, error)
	DeleteSession(sessionID string) error
}

type RedisSessionStorage struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisSessionStorage(client *redis.Client, ttl int) *RedisSessionStorage {
	return &RedisSessionStorage{
		client: client,
		ttl:    time.Duration(ttl) * time.Second,
	}
}

func (r *RedisSessionStorage) CreateSession(session *models.Session) error {
	ctx := context.Background()
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("error marshalling session: %v", err)
	}
	status := r.client.Set(ctx, session.ID, data, r.ttl)
	if err := status.Err(); err != nil {
		return fmt.Errorf("error saving session to redis: %v", err)
	}
	return nil
}

func (r *RedisSessionStorage) UpdateSession(session *models.Session) error {
	ctx := context.Background()
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("error marshalling session: %v", err)
	}

	status := r.client.Set(ctx, session.ID, data, r.ttl)
	if err := status.Err(); err != nil {
		return fmt.Errorf("error saving session to redis: %v", err)
	}
	return nil
}

func (r *RedisSessionStorage) GetSession(sessionID string) (*models.Session, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, sessionID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("session not found")
	} else if err != nil {
		return nil, fmt.Errorf("error retrieving session from redis: %v", err)
	}

	var session models.Session
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, fmt.Errorf("error unmarshalling session data: %v", err)
	}

	status := r.client.Expire(ctx, sessionID, r.ttl)
	if err := status.Err(); err != nil {
		return nil, fmt.Errorf("error resetting TTL for session: %v", err)
	}

	return &session, nil
}

func (r *RedisSessionStorage) DeleteSession(sessionID string) error {
	ctx := context.Background()
	_, err := r.client.Del(ctx, sessionID).Result()
	if err != nil {
		return fmt.Errorf("error deleting session from redis: %v", err)
	}
	return nil
}

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
)

func createKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

type SessionRepository struct {
	RedisClient *redis.Client
}

func loadRedisConfiguration(options map[string]interface{}) (*redis.Options, error) {
	addr, ok := options["addr"].(string)
	if !ok {
		return nil, errors.New("invalid type for redis addr")
	}

	password, ok := options["password"].(string)
	if !ok {
		return nil, errors.New("invalid type for redis password")
	}

	db, ok := options["db"].(int)
	if !ok {
		return nil, errors.New("invalid type for redis db")
	}

	return &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}, nil
}

func (st *SessionRepository) PutSession(s domain.Session) error {
	serializedSession, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = st.RedisClient.Set(
		createKey(s.ID),
		serializedSession,
		s.ExpiryTime.Sub(time.Now().UTC()),
	).Err()
	if err != nil {
		return err
	}

	return nil
}

func (st *SessionRepository) GetSessionByID(id string) (*domain.Session, error) {

	serializedSession, err := st.RedisClient.Get(createKey(id)).Bytes()
	if err != nil {
		return nil, err
	}

	var session domain.Session
	if err := json.Unmarshal(serializedSession, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (st *SessionRepository) DeleteSession(s domain.Session) error {
	return st.RedisClient.Del(createKey(s.ID)).Err()
}

func NewSessionRepository(options map[string]interface{}) (*SessionRepository, error) {
	redisOpts, err := loadRedisConfiguration(options)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(redisOpts)

	_, err = redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &SessionRepository{
		RedisClient: redisClient,
	}, nil
}

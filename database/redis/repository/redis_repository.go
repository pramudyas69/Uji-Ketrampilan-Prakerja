package repository

import (
	"context"
	"fmt"
	"time"
	"uji/database/redis"
)

type RedisRepository interface {
	SetValue(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetValue(ctx context.Context, key string) (string, error)
	DeleteKey(ctx context.Context, key string) error
}

type redisRepository struct {
	redisClient *redis.RedisClient
}

func NewRedisRepository(redisClient *redis.RedisClient) RedisRepository {
	return &redisRepository{redisClient: redisClient}
}

func (r *redisRepository) SetValue(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.redisClient.SetValue(ctx, key, value, expiration)
	if err != nil {
		return fmt.Errorf("failed to set value in Redis repository: %v", err)
	}
	return nil
}

func (r *redisRepository) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.GetValue(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get value from Redis repository: %v", err)
	}
	return val, nil
}

func (r *redisRepository) DeleteKey(ctx context.Context, key string) error {
	err := r.redisClient.DeleteKey(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete key from Redis repository: %v", err)
	}
	return nil
}

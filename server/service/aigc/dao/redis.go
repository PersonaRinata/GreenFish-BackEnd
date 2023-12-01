package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisManager struct {
	redisClient *redis.Client
}

func NewRedisManager(client *redis.Client) *RedisManager {
	return &RedisManager{redisClient: client}
}

func (r *RedisManager) GetAIGCHistory(ctx context.Context, userID int64) ([]string, error) {
	res, err := r.redisClient.LRange(ctx, "aigc:"+strconv.FormatInt(userID, 10), 0, -1).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

func (r *RedisManager) AddAIGCRecord(ctx context.Context, userID int64, msg ...string) error {
	err := r.redisClient.LPush(ctx, "aigc:"+strconv.FormatInt(userID, 10), msg).Err()
	if err != nil {
		return err
	}
	return nil
}

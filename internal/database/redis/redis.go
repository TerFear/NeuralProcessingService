package redis

import (
	"context"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type Redis struct {
	Redis *redis.Client
	log   *zap.Logger
}

func New(log *zap.Logger) *Redis {
	r := &Redis{
		log: log,
		Redis: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	err := r.Ping()

	if err != nil {
		log.Fatal("Redis connect failed", zap.Error(err))
	}

	return r
}

func (r *Redis) Ping() (err error) {
	err = r.Redis.Ping().Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) CheckHash(ctx context.Context, answer string) (data string, err error) {
	res, err := r.Redis.HGetAll(answer).Result()

	if err != nil {
		return "", err
	}

	data = res["data"]

	return data, nil
}

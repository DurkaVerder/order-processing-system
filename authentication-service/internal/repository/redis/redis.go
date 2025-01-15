package redis

import (
	"os"

	"github.com/go-redis/redis"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	return &Redis{
		rdb: rdb,
	}
}

func (r *Redis) RevokeToken(token string) error {
	err := r.rdb.Set(token, "revoked", 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) IsTokenRevoked(token string) (bool, error) {
	val, err := r.rdb.Get(token).Result()
	if err != nil {
		return false, err
	}

	if val == "" {
		return false, nil
	}

	return true, nil
}

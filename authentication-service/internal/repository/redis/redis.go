package redis

import (
	"os"

	"github.com/go-redis/redis"
)

// Redis is a struct that contains a pointer to a redis client
type Redis struct {
	rdb *redis.Client
}

// NewRedis creates a new Redis struct and returns a pointer to it
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

// RevokeToken sets a token as revoked in the redis database
func (r *Redis) RevokeToken(token string) error {
	err := r.rdb.Set(token, "revoked", 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// IsTokenRevoked checks if a token is revoked
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

package repository

import "github.com/go-redis/redis"

type redisRepository struct {
	client *redis.Client
}

// NewRedisRepository creates a new RedisRepository
func NewRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Get(identifier string) (int, error) {
	return r.client.Get(identifier).Int()
}

func (r *redisRepository) Incr(identifier string) error {
	return r.client.Incr(identifier).Err()
}

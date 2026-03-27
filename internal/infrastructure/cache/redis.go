package cache

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisCache[V any] struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache[V any](Password string) *RedisCache[V] {
	return &RedisCache[V]{
		client: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: Password,
				DB:       0,
			},
		),
		ctx: context.Background(),
	}
}

func (r *RedisCache[V]) Get(key string) (V, bool) {

	var result V

	val, err := r.client.Get(r.ctx, key).Result()

	if err != nil {
		return result, false
	}

	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return result, false
	}

	return result, true
}

func (r *RedisCache[V]) Set(key string, value V) error {

	data, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, key, data, 0).Err()
}

func (r *RedisCache[V]) Update(key string, value V) error {

	exists, err := r.client.Exists(r.ctx, key).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return errors.New("key does not exist")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, key, data, 0).Err()

}

func (r *RedisCache[V]) Delete(key string) error {

	return r.client.Del(r.ctx, key).Err()

}

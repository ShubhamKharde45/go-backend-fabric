package service

import (
	"github.com/ShubhamKharde45/rate_limiter/internal/domain"
)

type CacheService[V any] struct {
	Store domain.Cache[V]
}

func NewCacheService[K comparable, V any](store domain.Cache[V]) *CacheService[V] {
	return &CacheService[V]{
		Store: store,
	}
}

func (c *CacheService[V]) Get(Key string) (V, bool) {
	value, ok := c.Store.Get(Key)

	if !ok {
		return value, false
	}

	return value, true

}

func (c *CacheService[V]) Set(Key string, Value V) error {
	err := c.Store.Set(Key, Value)

	if err != nil {
		return err
	}

	return nil

}

func (c *CacheService[V]) Update(Key string, Value V) error {
	err := c.Store.Update(Key, Value)

	if err != nil {
		return err
	}

	return nil

}

func (c *CacheService[V]) Delete(Key string) error {
	err := c.Store.Delete(Key)

	if err != nil {
		return err
	}

	return nil

}

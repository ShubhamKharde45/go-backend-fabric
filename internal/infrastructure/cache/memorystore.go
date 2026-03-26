package cache

import (
	"sync"

	"github.com/ShubhamKharde45/rate_limiter/internal/domain"
)

type MemoryStore[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func NewMemoryStore[K comparable, V any]() *MemoryStore[K, V] {
	return &MemoryStore[K, V]{
		data: make(map[K]V),
	}
}

func (ms *MemoryStore[K, V]) Get(key K) (V, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	value, ok := ms.data[key]
	return value, ok
}

func (ms *MemoryStore[K, V]) Set(key K, value V) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.data[key]; ok {
		return domain.ErrAlreadyExist
	}

	ms.data[key] = value
	return nil
}

func (ms *MemoryStore[K, V]) Update(key K, value V) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.data[key]; !ok {
		return domain.ErrNotExist
	}

	ms.data[key] = value
	return nil
}

func (ms *MemoryStore[K, V]) Delete(key K) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.data[key]; !ok {
		return domain.ErrNotExist
	}

	delete(ms.data, key)
	return nil
}

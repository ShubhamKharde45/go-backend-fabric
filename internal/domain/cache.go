package domain

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V) error

	Update(key K, value V) error
	Delete(key K) error
}

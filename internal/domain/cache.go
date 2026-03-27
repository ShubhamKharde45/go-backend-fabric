package domain

type Cache[V any] interface {
	Get(key string) (V, bool)
	Set(key string, value V) error

	Update(key string, value V) error
	Delete(key string) error
}

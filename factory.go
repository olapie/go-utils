package utils

import (
	"context"
	"sync"
)

type NamedInstanceInitializer[T any] func(ctx context.Context, key string) (T, error)

type NamedInstanceFactory[T any] struct {
	mu          sync.RWMutex // one lock is okay, as most cases are just read lock
	cache       map[string]T
	initializer NamedInstanceInitializer[T]
}

func NewNamedInstanceFactory[T any](initializer NamedInstanceInitializer[T]) *NamedInstanceFactory[T] {
	return &NamedInstanceFactory[T]{
		cache:       map[string]T{},
		initializer: initializer,
	}
}

func (f *NamedInstanceFactory[T]) Get(ctx context.Context, name string) (instance T, err error) {
	f.mu.RLock()
	instance, ok := f.cache[name]
	f.mu.RUnlock()
	if ok {
		return instance, nil
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	instance, ok = f.cache[name]
	if !ok {
		instance, err = f.initializer(ctx, name)
		if err != nil {
			return instance, err
		}
		f.cache[name] = instance
	}
	return instance, nil
}

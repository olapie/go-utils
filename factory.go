package utils

import (
	"context"
	"sync"
)

type NamedInstanceInitializer[T any] func(ctx context.Context, key string) T

type NamedInstanceFactory[T any] struct {
	mu          sync.RWMutex // one lock is okay, as most cases are just read lock
	cache       map[string]T
	initializer NamedInstanceInitializer[T]
}

func NewNamedInstanceFactory[T any](initializer NamedInstanceInitializer[T]) *Factory[T] {
	return &NamedInstanceFactory[T]{
		cache:       map[string]T{},
		initializer: initializer,
	}
}

func (f *NamedInstanceFactory[T]) Get(ctx context.Context, name string) T {
	f.mu.RLock()
	instance, ok := f.cache[name]
	f.mu.RUnlock()
	if ok {
		return instance
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	instance, ok = f.cache[name]
	if !ok {
		instance = f.initializer(ctx, name)
		f.cache[name] = instance
	}
	return instance
}

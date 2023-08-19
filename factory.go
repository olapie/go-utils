package utils

import (
	"context"
	"sync"
)

type Factory[T any] struct {
	mu      sync.RWMutex // one lock is okay, as most cases are just read lock
	cache   map[string]T
	creator func() T
}

func NewFactory[T any](creator func() T) *Factory[T] {
	return &Factory[T]{
		cache:   map[string]T{},
		creator: creator,
	}
}

func (f *Factory[T]) Get(ctx context.Context, key string) T {
	f.mu.RLock()
	instance, ok := f.cache[key]
	f.mu.RUnlock()
	if ok {
		return instance
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	instance, ok = f.cache[key]
	if !ok {
		instance = f.creator()
		f.cache[key] = instance
	}
	return instance
}

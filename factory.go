package utils

import (
	"context"
	"sync"
)

type ObjectFactory[T any] struct {
	mu          sync.RWMutex // one lock is okay, as most cases are just read lock
	cache       map[string]T
	initializer func(ctx context.Context, id string) T
}

func NewObjectFactory[T any](initializer func(ctx context.Context, id string) T) *ObjectFactory[T] {
	return &ObjectFactory[T]{
		cache:       map[string]T{},
		initializer: initializer,
	}
}

func (f *ObjectFactory[T]) Get(ctx context.Context, id string) T {
	f.mu.RLock()
	obj, ok := f.cache[id]
	f.mu.RUnlock()
	if ok {
		return obj
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	obj, ok = f.cache[id]
	if !ok {
		obj = f.initializer(ctx, id)
	}
	f.cache[id] = obj
	return obj
}

type ObjectFactoryE[T any] struct {
	mu          sync.RWMutex // one lock is okay, as most cases are just read lock
	cache       map[string]T
	initializer func(ctx context.Context, id string) (T, error)
}

func NewObjectFactoryE[T any](initializer func(ctx context.Context, id string) (T, error)) *ObjectFactoryE[T] {
	return &ObjectFactoryE[T]{
		cache:       map[string]T{},
		initializer: initializer,
	}
}

func (f *ObjectFactoryE[T]) Get(ctx context.Context, id string) (obj T, err error) {
	f.mu.RLock()
	obj, ok := f.cache[id]
	f.mu.RUnlock()
	if ok {
		return obj, nil
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	obj, ok = f.cache[id]
	if ok {
		return obj, nil
	}
	obj, err = f.initializer(ctx, id)
	if err != nil {
		return obj, err
	}
	f.cache[id] = obj
	return obj, nil
}

package queue

import "sync"

type mutexQueue[V any] struct {
	mu     sync.RWMutex
	values []V
}

func newMutexQueue[V any]() *mutexQueue[V] {
	return &mutexQueue[V]{mu: sync.RWMutex{}, values: make([]V, 0)}
}

func (q *mutexQueue[V]) Push(value V) {
	q.mu.Lock()
	q.values = append(q.values, value)
	q.mu.Unlock()
}

func (q *mutexQueue[V]) Pop() (value V, ok bool) {
	q.mu.Lock()
	if len(q.values) == 0 {
		q.mu.Unlock()
		return value, false
	}
	value, q.values = q.values[0], q.values[1:]
	q.mu.Unlock()
	return value, true
}

func (q *mutexQueue[V]) Len() int {
	return len(q.values)
}

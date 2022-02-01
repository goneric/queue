package queue

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

func TestQueue_PopEmpty(t *testing.T) {
	q := New[int]()
	_, ok := q.Pop()
	if ok {
		t.Fatalf("pop empty queue should return false")
	}
}

func TestQueue_Len(t *testing.T) {
	q := New[int]()
	if q.Len() != 0 {
		t.Fatalf("empty queue should have zero length")
	}
	q.Push(1)
	q.Push(2)
	if q.Len() != 2 {
		t.Fatalf("expect length = %d, got %d", 2, q.Len())
	}
	q.Pop()
	if q.Len() != 1 {
		t.Fatalf("expect length = %d, got %d", 1, q.Len())
	}
}

func TestConcurrentWrite(t *testing.T) {
	n := 1 << 5
	inputs := make([]string, n)
	for i := 0; i < n; i++ {
		input := fmt.Sprintf("input_%d", i)
		inputs = append(inputs, input)
	}

	var wg sync.WaitGroup
	q := New[string]()

	for i := 0; i < n; i++ {
		input := inputs[i]
		wg.Add(1)
		go func() {
			q.Push(input)
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkQueue(b *testing.B) {
	n := 1 << 10
	inputs := make([]int, n)
	for i := 0; i < n; i++ {
		inputs = append(inputs, rand.Int())
	}
	b.ResetTimer()

	// lfq := newLockFreeQueue[int]()
	q := newMutexQueue[int]()
	b.Run(fmt.Sprintf("%T", q), func(b *testing.B) {
		var c int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := int(atomic.AddInt64(&c, 1)-1) % n
				v := inputs[i]
				if v >= 0 {
					q.Push(v)
				} else {
					q.Pop()
				}
			}
		})
	})
}

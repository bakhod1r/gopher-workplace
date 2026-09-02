package observerbus

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestPublishReachesAll(t *testing.T) {
	b := NewBus()
	var a, c int64

	b.Subscribe(HandlerFunc(func(string) { atomic.AddInt64(&a, 1) }))
	b.Subscribe(HandlerFunc(func(string) { atomic.AddInt64(&c, 1) }))

	b.Publish("e")
	if a != 1 || c != 1 {
		t.Errorf("counts = %d, %d; want 1, 1", a, c)
	}
	if b.Count() != 2 {
		t.Errorf("Count = %d, want 2", b.Count())
	}
}

func TestUnsubscribe(t *testing.T) {
	b := NewBus()
	var n int64
	unsub := b.Subscribe(HandlerFunc(func(string) { atomic.AddInt64(&n, 1) }))

	b.Publish("e")
	unsub()
	b.Publish("e")

	if n != 1 {
		t.Errorf("handler ran %d times, want 1", n)
	}
	if b.Count() != 0 {
		t.Errorf("Count = %d, want 0", b.Count())
	}
}

func TestUnsubscribeTwiceIsSafe(t *testing.T) {
	b := NewBus()
	unsub := b.Subscribe(HandlerFunc(func(string) {}))
	unsub()
	unsub()
	if b.Count() != 0 {
		t.Errorf("Count = %d, want 0", b.Count())
	}
}

func TestUnsubscribeDuringPublish(t *testing.T) {
	b := NewBus()
	var n int64

	var unsub func()
	unsub = b.Subscribe(HandlerFunc(func(string) {
		atomic.AddInt64(&n, 1)
		unsub()
	}))
	b.Subscribe(HandlerFunc(func(string) {}))

	b.Publish("first")
	b.Publish("second")

	if n != 1 {
		t.Errorf("self-unsubscribing handler ran %d times, want 1", n)
	}
	if b.Count() != 1 {
		t.Errorf("Count = %d, want 1", b.Count())
	}
}

func TestSubscribeDuringPublish(t *testing.T) {
	b := NewBus()
	var added int64

	b.Subscribe(HandlerFunc(func(string) {
		if atomic.AddInt64(&added, 1) == 1 {
			b.Subscribe(HandlerFunc(func(string) {}))
		}
	}))

	b.Publish("e")
	if b.Count() != 2 {
		t.Errorf("Count = %d, want 2", b.Count())
	}
}

func TestConcurrentPublishAndSubscribe(t *testing.T) {
	b := NewBus()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			unsub := b.Subscribe(HandlerFunc(func(string) {}))
			unsub()
		}()
		go func() {
			defer wg.Done()
			b.Publish("e")
		}()
	}
	wg.Wait()

	if b.Count() != 0 {
		t.Errorf("Count = %d, want 0", b.Count())
	}
}

func TestPublishWithNoSubscribers(t *testing.T) {
	NewBus().Publish("e")
}

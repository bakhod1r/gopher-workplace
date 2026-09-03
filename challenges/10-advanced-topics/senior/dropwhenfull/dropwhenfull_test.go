package dropwhenfull

import (
	"sync"
	"testing"
	"time"
)

func TestOfferWithRoom(t *testing.T) {
	ch := make(chan int, 2)
	if !Offer(ch, 1) {
		t.Error("Offer = false, want true: the buffer had room")
	}
	if got := <-ch; got != 1 {
		t.Errorf("received %d, want 1", got)
	}
}

func TestOfferWhenFull(t *testing.T) {
	ch := make(chan int, 1)
	Offer(ch, 1)
	if Offer(ch, 2) {
		t.Error("Offer = true, want false: the buffer was full")
	}
	if got := <-ch; got != 1 {
		t.Errorf("received %d, want 1: the dropped value must not displace the first", got)
	}
}

func TestOfferDoesNotBlock(t *testing.T) {
	ch := make(chan int) // unbuffered, no receiver
	done := make(chan bool, 1)
	go func() { done <- Offer(ch, 1) }()
	select {
	case ok := <-done:
		if ok {
			t.Error("Offer = true on an unbuffered channel with no receiver")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Offer blocked, want an immediate false")
	}
}

func TestOfferUnbufferedWithReceiver(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	got := make(chan int, 1)
	go func() {
		defer wg.Done()
		got <- <-ch
	}()
	deadline := time.Now().Add(2 * time.Second)
	for !Offer(ch, 7) && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
	if v := <-got; v != 7 {
		t.Errorf("received %d, want 7", v)
	}
}

func TestOfferFillsExactly(t *testing.T) {
	ch := make(chan int, 4)
	accepted := 0
	for i := 0; i < 10; i++ {
		if Offer(ch, i) {
			accepted++
		}
	}
	if accepted != 4 {
		t.Errorf("accepted %d, want 4", accepted)
	}
}

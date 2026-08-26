package actorpatt

import (
	"sync"
	"testing"
)

func TestActor(t *testing.T) {
	a := New()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Add(1)
		}()
	}
	wg.Wait()

	// Wait for queue to process
	done := make(chan int)
	a.msgs <- func(val *int) { done <- *val }

	if got := <-done; got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

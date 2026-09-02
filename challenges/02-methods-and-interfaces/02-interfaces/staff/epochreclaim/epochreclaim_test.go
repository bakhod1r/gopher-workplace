package epochreclaim

import (
	"sync"
	"testing"
)

func TestReclaimWithNoReaders(t *testing.T) {
	e := NewEpoch()
	n := &Node{ID: 1}
	e.Retire(n)

	if e.Pending() != 1 {
		t.Errorf("Pending = %d, want 1", e.Pending())
	}
	if got := e.Reclaim(); got != 1 {
		t.Errorf("Reclaim = %d, want 1", got)
	}
	if !n.Freed {
		t.Error("node was not freed")
	}
	if e.Pending() != 0 {
		t.Errorf("Pending = %d, want 0", e.Pending())
	}
}

func TestReaderDelaysReclaim(t *testing.T) {
	e := NewEpoch()
	ep := e.Enter()

	n := &Node{ID: 1}
	e.Retire(n)

	if got := e.Reclaim(); got != 0 {
		t.Errorf("Reclaim = %d, want 0 while a reader is active", got)
	}
	if n.Freed {
		t.Fatal("freed while a reader could still be reading it")
	}

	e.Exit(ep)
	if got := e.Reclaim(); got != 1 {
		t.Errorf("Reclaim = %d, want 1 after the reader exited", got)
	}
	if !n.Freed {
		t.Error("node was not freed after the grace period")
	}
}

func TestLaterReaderDoesNotBlockReclaim(t *testing.T) {
	e := NewEpoch()

	n := &Node{ID: 1}
	e.Retire(n) // retired in epoch 0, current advances to 1

	ep := e.Enter() // enters epoch 1, cannot have seen the node
	if got := e.Reclaim(); got != 1 {
		t.Errorf("Reclaim = %d, want 1 (the new reader cannot see the node)", got)
	}
	e.Exit(ep)
}

func TestMultipleReaders(t *testing.T) {
	e := NewEpoch()
	a := e.Enter()
	b := e.Enter()

	n := &Node{ID: 1}
	e.Retire(n)

	e.Exit(a)
	if got := e.Reclaim(); got != 0 {
		t.Errorf("Reclaim = %d, want 0 while one reader remains", got)
	}

	e.Exit(b)
	if got := e.Reclaim(); got != 1 {
		t.Errorf("Reclaim = %d, want 1", got)
	}
}

func TestManyRetiredObjects(t *testing.T) {
	e := NewEpoch()
	nodes := make([]*Node, 10)
	for i := range nodes {
		nodes[i] = &Node{ID: i}
		e.Retire(nodes[i])
	}

	if got := e.Reclaim(); got != 10 {
		t.Errorf("Reclaim = %d, want 10", got)
	}
	for _, n := range nodes {
		if !n.Freed {
			t.Fatalf("node %d was not freed", n.ID)
		}
	}
}

func TestConcurrentReadersAndRetires(t *testing.T) {
	e := NewEpoch()
	var wg sync.WaitGroup

	for r := 0; r < 8; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				ep := e.Enter()
				e.Exit(ep)
			}
		}()
	}

	nodes := make([]*Node, 200)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range nodes {
			nodes[i] = &Node{ID: i}
			e.Retire(nodes[i])
			e.Reclaim()
		}
	}()

	wg.Wait()

	// With every reader finished, a final Reclaim must free the rest.
	e.Reclaim()
	if e.Pending() != 0 {
		t.Errorf("Pending = %d, want 0 once all readers have exited", e.Pending())
	}
	for _, n := range nodes {
		if !n.Freed {
			t.Fatalf("node %d leaked", n.ID)
		}
	}
}

func TestIsFreer(t *testing.T) {
	var f Freer = &Node{ID: 1}
	f.Free()
	if !f.(*Node).Freed {
		t.Error("Free did not mark the node")
	}
}

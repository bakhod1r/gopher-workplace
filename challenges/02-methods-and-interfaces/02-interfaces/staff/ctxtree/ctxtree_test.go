package ctxtree

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func isClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func TestRootStartsLive(t *testing.T) {
	r := NewRoot()
	if isClosed(r.Done()) {
		t.Error("a new root should not be cancelled")
	}
	if r.Err() != nil {
		t.Errorf("Err = %v, want nil", r.Err())
	}
}

func TestCancelPropagatesDown(t *testing.T) {
	root := NewRoot()
	a := root.Child()
	b := a.Child()
	c := root.Child()

	root.Cancel()

	for name, n := range map[string]*Node{"root": root, "a": a, "b": b, "c": c} {
		if !isClosed(n.Done()) {
			t.Errorf("%s was not cancelled", name)
		}
		if !errors.Is(n.Err(), ErrCancelled) {
			t.Errorf("%s: Err = %v, want ErrCancelled", name, n.Err())
		}
	}
}

func TestCancelChildLeavesSiblings(t *testing.T) {
	root := NewRoot()
	a := root.Child()
	b := root.Child()

	a.Cancel()

	if !isClosed(a.Done()) {
		t.Error("a should be cancelled")
	}
	if isClosed(b.Done()) {
		t.Error("b is a sibling and must stay live")
	}
	if isClosed(root.Done()) {
		t.Error("cancelling a child must not cancel the parent")
	}
}

func TestChildOfCancelledIsBornCancelled(t *testing.T) {
	root := NewRoot()
	root.Cancel()

	c := root.Child()
	if !isClosed(c.Done()) {
		t.Error("a child of a cancelled node should be born cancelled")
	}
	if !errors.Is(c.Err(), ErrCancelled) {
		t.Errorf("Err = %v, want ErrCancelled", c.Err())
	}
}

func TestCancelIsIdempotent(t *testing.T) {
	root := NewRoot()
	child := root.Child()

	root.Cancel()
	root.Cancel()
	child.Cancel()

	if !isClosed(root.Done()) || !isClosed(child.Done()) {
		t.Error("nodes should stay cancelled")
	}
}

func TestWaitersAreWoken(t *testing.T) {
	root := NewRoot()
	child := root.Child()

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-child.Done():
			case <-time.After(5 * time.Second):
				t.Error("waiter was never woken")
			}
		}()
	}

	root.Cancel()
	wg.Wait()
}

func TestConcurrentCancelAndChild(t *testing.T) {
	root := NewRoot()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := root.Child()
			c.Child()
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		root.Cancel()
	}()
	wg.Wait()

	// After everything settles, one more cancel must be safe and the root
	// must report cancelled.
	root.Cancel()
	if !errors.Is(root.Err(), ErrCancelled) {
		t.Errorf("Err = %v, want ErrCancelled", root.Err())
	}
}

func TestIsCanceller(t *testing.T) {
	var c Canceller = NewRoot()
	c.Cancel()
	if !errors.Is(c.Err(), ErrCancelled) {
		t.Errorf("Err = %v, want ErrCancelled", c.Err())
	}
}

func TestDeepChain(t *testing.T) {
	root := NewRoot()
	cur := root
	nodes := []*Node{root}
	for i := 0; i < 100; i++ {
		cur = cur.Child()
		nodes = append(nodes, cur)
	}

	root.Cancel()
	for i, n := range nodes {
		if !isClosed(n.Done()) {
			t.Fatalf("node %d was not cancelled", i)
		}
	}
}

package dsucompressbug

import (
	"testing"
	"time"
)

func TestChainedUnions(t *testing.T) {
	u := &DSU[int]{}
	u.Union(1, 2)
	u.Union(2, 3)
	if !u.Connected(1, 3) {
		t.Error("Connected(1,3) = false, want true")
	}
}

func TestUnseenElement(t *testing.T) {
	u := &DSU[int]{}
	u.Union(1, 2)
	if u.Connected(1, 9) {
		t.Error("Connected(1,9) = true, want false")
	}
}

func TestDegenerateChainCost(t *testing.T) {
	const n = 15000
	u := &DSU[int]{}
	for i := 0; i < n; i++ {
		u.Union(i, i+1)
	}
	start := time.Now()
	for i := 0; i <= n; i++ {
		if !u.Connected(0, i) {
			t.Fatalf("Connected(0,%d) = false, want true", i)
		}
	}
	if d := time.Since(start); d > 300*time.Millisecond {
		t.Fatalf("%d queries over a %d-deep chain took %v, want under 300ms", n, n, d)
	}
}

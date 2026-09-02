package sparsegridkeybug

import (
	"testing"
	"time"
)

func TestGridSwappedCoordinates(t *testing.T) {
	var g Grid[string]
	g.Set(Point{X: 1, Y: 2}, "a")
	g.Set(Point{X: 2, Y: 1}, "b")
	if g.Len() != 2 {
		t.Fatalf("Len = %d, want 2", g.Len())
	}
	if v, ok := g.Get(Point{X: 1, Y: 2}); !ok || v != "a" {
		t.Errorf("Get({1,2}) = %q, %v, want a, true", v, ok)
	}
	if v, ok := g.Get(Point{X: 2, Y: 1}); !ok || v != "b" {
		t.Errorf("Get({2,1}) = %q, %v, want b, true", v, ok)
	}
}

func TestGridNegativeCoordinates(t *testing.T) {
	var g Grid[string]
	g.Set(Point{X: -1, Y: 5}, "a")
	g.Set(Point{X: 4, Y: 0}, "b")
	if g.Len() != 2 {
		t.Fatalf("Len = %d, want 2", g.Len())
	}
	if v, ok := g.Get(Point{X: -1, Y: 5}); !ok || v != "a" {
		t.Errorf("Get({-1,5}) = %q, %v, want a, true", v, ok)
	}
	if _, ok := g.Get(Point{X: -2, Y: 5}); ok {
		t.Error("Get({-2,5}) = ok, want false")
	}
}

func TestGridDenseBlock(t *testing.T) {
	const n = 200
	start := time.Now()
	var g Grid[int]
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			g.Set(Point{X: x, Y: y}, x*n+y)
		}
	}
	if g.Len() != n*n {
		t.Fatalf("Len = %d, want %d", g.Len(), n*n)
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v, ok := g.Get(Point{X: x, Y: y})
			if !ok || v != x*n+y {
				t.Fatalf("Get({%d,%d}) = %d, %v, want %d, true", x, y, v, ok, x*n+y)
			}
		}
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("40k cells took %v, want under 2s", d)
	}
}

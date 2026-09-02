package workstealq

import (
	"sync"
	"testing"
)

func TestOwnerIsLIFO(t *testing.T) {
	d := NewDeque()
	d.Push(1)
	d.Push(2)

	if v, ok := d.Pop(); v != 2 || !ok {
		t.Errorf("Pop = %d, %v; want 2, true", v, ok)
	}
	if v, ok := d.Pop(); v != 1 || !ok {
		t.Errorf("Pop = %d, %v; want 1, true", v, ok)
	}
	if _, ok := d.Pop(); ok {
		t.Error("Pop on an empty deque returned ok")
	}
}

func TestThiefIsFIFO(t *testing.T) {
	d := NewDeque()
	d.Push(1)
	d.Push(2)
	d.Push(3)

	if v, ok := d.Steal(); v != 1 || !ok {
		t.Errorf("Steal = %d, %v; want 1, true", v, ok)
	}
	if v, ok := d.Steal(); v != 2 || !ok {
		t.Errorf("Steal = %d, %v; want 2, true", v, ok)
	}
	if v, ok := d.Pop(); v != 3 || !ok {
		t.Errorf("Pop = %d, %v; want 3, true", v, ok)
	}
}

func TestEmptyDeque(t *testing.T) {
	d := NewDeque()
	if _, ok := d.Pop(); ok {
		t.Error("Pop on empty returned ok")
	}
	if _, ok := d.Steal(); ok {
		t.Error("Steal on empty returned ok")
	}
	if d.Len() != 0 {
		t.Errorf("Len = %d, want 0", d.Len())
	}
}

func TestLen(t *testing.T) {
	d := NewDeque()
	d.Push(1)
	d.Push(2)
	if d.Len() != 2 {
		t.Errorf("Len = %d, want 2", d.Len())
	}
	d.Steal()
	if d.Len() != 1 {
		t.Errorf("Len = %d, want 1", d.Len())
	}
	d.Pop()
	if d.Len() != 0 {
		t.Errorf("Len = %d, want 0", d.Len())
	}
}

func TestNoDuplicatesUnderContention(t *testing.T) {
	const n = 5000
	d := NewDeque()
	for i := 0; i < n; i++ {
		d.Push(i)
	}

	var mu sync.Mutex
	seen := make(map[int]int, n)
	record := func(v int) {
		mu.Lock()
		seen[v]++
		mu.Unlock()
	}

	var wg sync.WaitGroup

	// The owner pops from the bottom.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := d.Pop()
			if !ok {
				return
			}
			record(v)
		}
	}()

	// Thieves steal from the top.
	for th := 0; th < 4; th++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				v, ok := d.Steal()
				if !ok {
					if d.Len() == 0 {
						return
					}
					continue
				}
				record(v)
			}
		}()
	}
	wg.Wait()

	if len(seen) != n {
		t.Fatalf("saw %d distinct items, want %d", len(seen), n)
	}
	for v, c := range seen {
		if c != 1 {
			t.Fatalf("item %d taken %d times", v, c)
		}
	}
}

func TestRaceOnLastItem(t *testing.T) {
	for trial := 0; trial < 200; trial++ {
		d := NewDeque()
		d.Push(42)

		var mu sync.Mutex
		takes := 0

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			if _, ok := d.Pop(); ok {
				mu.Lock()
				takes++
				mu.Unlock()
			}
		}()
		go func() {
			defer wg.Done()
			if _, ok := d.Steal(); ok {
				mu.Lock()
				takes++
				mu.Unlock()
			}
		}()
		wg.Wait()

		if takes != 1 {
			t.Fatalf("the last item was taken %d times, want exactly 1", takes)
		}
	}
}

func TestIsStealer(t *testing.T) {
	var s Stealer = NewDeque()
	if _, ok := s.Steal(); ok {
		t.Error("Steal on empty returned ok")
	}
}

func BenchmarkPushPop(b *testing.B) {
	d := NewDeque()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		d.Push(i)
		d.Pop()
	}
}

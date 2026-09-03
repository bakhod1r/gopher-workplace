package striped

import (
	"fmt"
	"sync"
	"testing"
)

func TestSetGet(t *testing.T) {
	m := New(8)
	m.Set("a", 1)
	m.Set("b", 2)
	if v, ok := m.Get("a"); v != 1 || !ok {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := m.Get("zz"); v != 0 || ok {
		t.Errorf("Get(zz) = %d, %v, want 0, false", v, ok)
	}
	if got := m.Len(); got != 2 {
		t.Errorf("Len = %d, want 2", got)
	}
}

func TestOverwrite(t *testing.T) {
	m := New(4)
	m.Set("a", 1)
	m.Set("a", 9)
	if v, _ := m.Get("a"); v != 9 {
		t.Errorf("Get(a) = %d, want 9", v)
	}
	if got := m.Len(); got != 1 {
		t.Errorf("Len = %d, want 1", got)
	}
}

func TestShardOfIsStable(t *testing.T) {
	m := New(16)
	for _, key := range []string{"", "a", "some-longer-key", "zzz"} {
		first := m.ShardOf(key)
		for i := 0; i < 100; i++ {
			if got := m.ShardOf(key); got != first {
				t.Fatalf("ShardOf(%q) returned %d then %d; it must be stable", key, first, got)
			}
		}
		if first < 0 || first >= m.Shards() {
			t.Fatalf("ShardOf(%q) = %d, out of range for %d shards", key, first, m.Shards())
		}
	}
}

func TestKeysSpreadAcrossShards(t *testing.T) {
	m := New(8)
	used := make(map[int]bool)
	for i := 0; i < 1000; i++ {
		used[m.ShardOf(fmt.Sprintf("key-%d", i))] = true
	}
	if len(used) < 6 {
		t.Errorf("1000 keys landed in only %d of 8 shards; the hash is not spreading them", len(used))
	}
}

func TestNonPositiveShardCount(t *testing.T) {
	for _, n := range []int{0, -4} {
		m := New(n)
		if m.Shards() != 1 {
			t.Errorf("New(%d).Shards() = %d, want 1", n, m.Shards())
		}
		m.Set("a", 1)
		if v, ok := m.Get("a"); v != 1 || !ok {
			t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
		}
	}
}

func TestConcurrentUse(t *testing.T) {
	m := New(16)
	const writers, each = 25, 400
	var wg sync.WaitGroup
	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < each; j++ {
				key := fmt.Sprintf("k-%d-%d", w, j)
				m.Set(key, j)
				if v, ok := m.Get(key); !ok || v != j {
					t.Errorf("Get(%q) = %d, %v, want %d, true", key, v, ok, j)
					return
				}
			}
		}()
	}
	wg.Wait()
	if got := m.Len(); got != writers*each {
		t.Errorf("Len = %d, want %d", got, writers*each)
	}
}

package memokeynormbug

import (
	"testing"
	"time"
)

func TestGetNormalisesKeys(t *testing.T) {
	var m Memo[int]
	calls := 0
	fn := func(Key) int { calls++; return 1 }
	m.Get(Key{"API.example.com", 443}, fn)
	m.Get(Key{"api.example.com", 443}, fn)
	m.Get(Key{" api.example.com ", 443}, fn)
	m.Get(Key{"Api.Example.Com", 443}, fn)
	if calls != 1 {
		t.Errorf("fn called %d times, want 1", calls)
	}
}

func TestGetSeparatesRealKeys(t *testing.T) {
	var m Memo[int]
	calls := 0
	fn := func(Key) int { calls++; return calls }
	m.Get(Key{"a", 1}, fn)
	m.Get(Key{"a", 2}, fn)
	m.Get(Key{"b", 1}, fn)
	if calls != 3 {
		t.Errorf("fn called %d times, want 3", calls)
	}
}

func TestGetPassesNormalisedKey(t *testing.T) {
	var m Memo[string]
	got := m.Get(Key{"  MiXeD.Host  ", 80}, func(k Key) string { return k.Host })
	if got != "mixed.host" {
		t.Errorf("fn saw host %q, want \"mixed.host\"", got)
	}
}

func TestGetScale(t *testing.T) {
	const n = 100_000
	spellings := []string{"API.example.com", "api.example.com", " Api.Example.Com ", "API.EXAMPLE.COM"}
	var m Memo[int]
	fn := func(Key) int {
		acc := 0
		for i := 0; i < 5000; i++ {
			acc += i
		}
		return acc
	}
	start := time.Now()
	for i := 0; i < n; i++ {
		m.Get(Key{spellings[i%len(spellings)], 443}, fn)
	}
	elapsed := time.Since(start)
	if elapsed > 250*time.Millisecond {
		t.Errorf("%d cached lookups took %v, want under 250ms", n, elapsed)
	}
}

package zerocopykeys

import (
	"testing"
)

func TestCount(t *testing.T) {
	m := map[string]int{}
	Count(m, [][]byte{[]byte("a"), []byte("b"), []byte("a")})
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("m = %v, want map[a:2 b:1]", m)
	}
}

func TestCountSkipsEmptyKeys(t *testing.T) {
	m := map[string]int{}
	Count(m, [][]byte{nil, {}, []byte("x")})
	if len(m) != 1 || m["x"] != 1 {
		t.Errorf("m = %v, want map[x:1]", m)
	}
}

func TestCountKeysSurviveBufferReuse(t *testing.T) {
	m := map[string]int{}
	buf := make([]byte, 4)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		Count(m, [][]byte{buf})
	}
	if len(m) != 26 {
		t.Fatalf("map has %d keys, want 26: the stored keys view the reused buffer", len(m))
	}
	for i := 0; i < 26; i++ {
		k := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if m[k] != 1 {
			t.Fatalf("m[%q] = %d, want 1", k, m[k])
		}
	}
}

func TestCountRepeatedKeysDoNotAllocate(t *testing.T) {
	m := map[string]int{}
	key := []byte("stable-key")
	Count(m, [][]byte{key})
	batch := [][]byte{key}
	if n := testing.AllocsPerRun(200, func() { Count(m, batch) }); n != 0 {
		t.Errorf("Count made %v allocations for an existing key, want 0: borrow the bytes for the lookup", n)
	}
	if m["stable-key"] < 200 {
		t.Errorf("counter = %d, want at least 200", m["stable-key"])
	}
}

func TestCountAccumulates(t *testing.T) {
	m := map[string]int{"a": 5}
	Count(m, [][]byte{[]byte("a")})
	if m["a"] != 6 {
		t.Errorf("m[a] = %d, want 6", m["a"])
	}
}

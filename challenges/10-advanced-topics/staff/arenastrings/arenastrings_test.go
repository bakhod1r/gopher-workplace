package arenastrings

import (
	"testing"
	"unsafe"
)

func TestIntern(t *testing.T) {
	got := Intern([]byte("abcd"), [][2]int{{0, 2}, {2, 4}})
	if len(got) != 2 || got[0] != "ab" || got[1] != "cd" {
		t.Errorf("Intern = %q, want [ab cd]", got)
	}
	if got := Intern(nil, nil); got != nil {
		t.Errorf("Intern = %q, want nil", got)
	}
}

func TestInternBadSpans(t *testing.T) {
	got := Intern([]byte("abcd"), [][2]int{{-1, 2}, {0, 99}, {3, 1}, {2, 2}})
	for i, s := range got {
		if s != "" {
			t.Errorf("span %d = %q, want empty", i, s)
		}
	}
}

func TestInternSurvivesArenaReuse(t *testing.T) {
	arena := make([]byte, 8)
	copy(arena, "firstrun")
	keys := Intern(arena, [][2]int{{0, 5}, {5, 8}})
	copy(arena, "OVERWRIT")
	if keys[0] != "first" || keys[1] != "run" {
		t.Errorf("keys = %q, want [first run]: the strings view the reused arena", keys)
	}
}

func TestInternKeysStayValidAsMapKeys(t *testing.T) {
	arena := make([]byte, 4)
	m := make(map[string]int)
	for i := 0; i < 26; i++ {
		for j := range arena {
			arena[j] = byte('a' + i)
		}
		for _, k := range Intern(arena, [][2]int{{0, 4}}) {
			m[k] = i
		}
	}
	if len(m) != 26 {
		t.Fatalf("map has %d keys, want 26: the keys changed after insertion", len(m))
	}
	for i := 0; i < 26; i++ {
		k := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if got, ok := m[k]; !ok || got != i {
			t.Fatalf("m[%q] = %d, %v, want %d, true", k, got, ok, i)
		}
	}
}

func TestInternDoesNotAliasTheArena(t *testing.T) {
	arena := []byte("abcdef")
	got := Intern(arena, [][2]int{{0, 3}})
	if unsafe.StringData(got[0]) == unsafe.SliceData(arena) {
		t.Error("the result aliases the arena")
	}
}

func TestInternUsesOneBlock(t *testing.T) {
	arena := make([]byte, 256)
	spans := make([][2]int, 32)
	for i := range spans {
		spans[i] = [2]int{i * 8, i*8 + 8}
	}
	var sink []string
	n := testing.AllocsPerRun(50, func() { sink = Intern(arena, spans) })
	_ = sink
	if n > 3 {
		t.Errorf("Intern made %v allocations for 32 spans, want a handful: copy into one block", n)
	}
}

package repeatbytes

import (
	"bytes"
	"testing"
)

func TestRepeat(t *testing.T) {
	if got := Repeat([]byte("ab"), 3); !bytes.Equal(got, []byte("ababab")) {
		t.Errorf("Repeat = %q, want \"ababab\"", got)
	}
	if got := Repeat([]byte("x"), 0); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
	if got := Repeat(nil, 4); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
	if got := Repeat([]byte("ab"), -1); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
}

func TestRepeatIsIndependent(t *testing.T) {
	b := []byte("ab")
	out := Repeat(b, 2)
	b[0] = 'z'
	if out[0] != 'a' {
		t.Error("the result shares storage with the input")
	}
}

func TestRepeatAllocatesOnce(t *testing.T) {
	b := []byte("abcdefgh")
	if n := testing.AllocsPerRun(100, func() { _ = Repeat(b, 64) }); n > 1 {
		t.Errorf("Repeat made %v allocations, want 1", n)
	}
}

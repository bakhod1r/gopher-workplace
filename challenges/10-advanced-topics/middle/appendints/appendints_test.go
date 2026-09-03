package appendints

import (
	"bytes"
	"testing"
)

func TestAppendInts(t *testing.T) {
	if got := AppendInts(nil, []int{1, 2, 3}); !bytes.Equal(got, []byte("1 2 3")) {
		t.Errorf("AppendInts = %q, want \"1 2 3\"", got)
	}
	if got := AppendInts([]byte("x:"), []int{-4}); !bytes.Equal(got, []byte("x:-4")) {
		t.Errorf("AppendInts = %q, want \"x:-4\"", got)
	}
	if got := AppendInts([]byte("keep"), nil); !bytes.Equal(got, []byte("keep")) {
		t.Errorf("AppendInts = %q, want \"keep\"", got)
	}
}

func TestAppendIntsAllocatesNothingWithRoom(t *testing.T) {
	vals := make([]int, 32)
	for i := range vals {
		vals[i] = i * 7
	}
	dst := make([]byte, 0, 512)
	if n := testing.AllocsPerRun(100, func() { _ = AppendInts(dst[:0], vals) }); n != 0 {
		t.Errorf("AppendInts made %v allocations, want 0: render into dst", n)
	}
}

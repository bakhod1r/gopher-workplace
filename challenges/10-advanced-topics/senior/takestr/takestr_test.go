package takestr

import (
	"testing"
	"unsafe"
)

func TestTake(t *testing.T) {
	if got := Take([]byte("hello"), 2); got != "he" {
		t.Errorf("Take = %q, want \"he\"", got)
	}
	if got := Take([]byte("hi"), 9); got != "hi" {
		t.Errorf("Take = %q, want \"hi\"", got)
	}
	if got := Take(nil, 3); got != "" {
		t.Errorf("Take = %q, want empty", got)
	}
	if got := Take([]byte("hi"), 0); got != "" {
		t.Errorf("Take = %q, want empty", got)
	}
}

func TestTakeSurvivesBufferReuse(t *testing.T) {
	buf := make([]byte, 8)
	copy(buf, "first")
	s := Take(buf, 5)
	copy(buf, "SECOND")
	if s != "first" {
		t.Errorf("s = %q, want \"first\": the string is a view of the reused buffer", s)
	}
}

func TestTakeResultsAreIndependent(t *testing.T) {
	buf := make([]byte, 4)
	got := make([]string, 0, 26)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		got = append(got, Take(buf, 4))
	}
	for i, s := range got {
		want := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if s != want {
			t.Fatalf("result %d = %q, want %q", i, s, want)
		}
	}
}

func TestTakeStillAvoidsTheDoubleCopy(t *testing.T) {
	buf := make([]byte, 256)
	var sink string
	n := testing.AllocsPerRun(100, func() { sink = Take(buf, 256) })
	_ = sink
	if n > 1 {
		t.Errorf("Take made %v allocations, want 1: one copy, then a header over it", n)
	}
}

func TestTakeDoesNotAliasTheInput(t *testing.T) {
	buf := []byte("abcd")
	s := Take(buf, 4)
	if unsafe.StringData(s) == unsafe.SliceData(buf) {
		t.Error("the result shares the caller's buffer")
	}
}

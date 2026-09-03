package interncache

import (
	"testing"
	"unsafe"
)

var sink string

func TestInternReturnsTheContents(t *testing.T) {
	var p Pool
	if got := p.Intern([]byte("hello")); got != "hello" {
		t.Errorf("Intern = %q, want \"hello\"", got)
	}
	if got := p.Intern(nil); got != "" {
		t.Errorf("Intern(nil) = %q, want empty", got)
	}
	if p.Len() != 1 {
		t.Errorf("Len = %d, want 1: the empty string is not stored", p.Len())
	}
}

func TestInternSharesOneCopy(t *testing.T) {
	var p Pool
	a := p.Intern([]byte("repeated"))
	b := p.Intern([]byte("repeated"))
	if unsafe.StringData(a) != unsafe.StringData(b) {
		t.Error("the two results are separate allocations; they must share one")
	}
	if p.Len() != 1 {
		t.Errorf("Len = %d, want 1", p.Len())
	}
}

func TestInternDistinctValues(t *testing.T) {
	var p Pool
	for i := 0; i < 26; i++ {
		p.Intern([]byte{byte('a' + i)})
	}
	if p.Len() != 26 {
		t.Errorf("Len = %d, want 26", p.Len())
	}
}

func TestInternSurvivesBufferReuse(t *testing.T) {
	var p Pool
	buf := make([]byte, 4)
	got := make([]string, 0, 26)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		got = append(got, p.Intern(buf))
	}
	for i, s := range got {
		want := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if s != want {
			t.Fatalf("result %d = %q, want %q: the stored string viewed the reused buffer", i, s, want)
		}
	}
	if p.Len() != 26 {
		t.Errorf("Len = %d, want 26", p.Len())
	}
}

func TestInternRepeatDoesNotAllocate(t *testing.T) {
	var p Pool
	key := []byte("a-repeated-token")
	p.Intern(key)
	if n := testing.AllocsPerRun(200, func() { sink = p.Intern(key) }); n != 0 {
		t.Errorf("a repeat Intern made %v allocations, want 0: borrow the bytes for the lookup", n)
	}
}

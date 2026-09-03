package valuebytes

import (
	"testing"
	"unsafe"
)

var sink []byte

func TestBytesShape(t *testing.T) {
	v := uint64(0)
	b := Bytes(&v)
	if len(b) != 8 {
		t.Errorf("len = %d, want 8", len(b))
	}
	if cap(b) != 8 {
		t.Errorf("cap = %d, want 8: an append must not run past the value", cap(b))
	}
}

func TestBytesSharesStorage(t *testing.T) {
	v := uint64(0)
	b := Bytes(&v)
	v = ^uint64(0)
	for i, x := range b {
		if x != 0xff {
			t.Fatalf("b[%d] = %#x, want 0xff: the view does not share v", i, x)
		}
	}
	b[0] = 0
	if v == ^uint64(0) {
		t.Error("writing through the view did not change v")
	}
}

func TestBytesNil(t *testing.T) {
	if got := Bytes(nil); got != nil {
		t.Errorf("Bytes(nil) = %v, want nil", got)
	}
}

func TestBytesAllocatesNothing(t *testing.T) {
	v := uint64(7)
	if n := testing.AllocsPerRun(200, func() { sink = Bytes(&v) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
}

func TestBytesMatchesTheSize(t *testing.T) {
	v := uint64(0)
	if uintptr(len(Bytes(&v))) != unsafe.Sizeof(v) {
		t.Error("the view's length must come from unsafe.Sizeof")
	}
}

package structbytes

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestBytesShape(t *testing.T) {
	f := &Frame{Kind: 1, Seq: 2, Stamp: 3}
	b, ok := Bytes(f)
	if !ok {
		t.Fatal("Bytes reported false for a pointer-free struct")
	}
	if uintptr(len(b)) != unsafe.Sizeof(*f) {
		t.Errorf("len = %d, want %d", len(b), unsafe.Sizeof(*f))
	}
	if uintptr(cap(b)) != unsafe.Sizeof(*f) {
		t.Errorf("cap = %d, want %d: an append must not run past the struct", cap(b), unsafe.Sizeof(*f))
	}
}

func TestBytesSharesTheStruct(t *testing.T) {
	f := &Frame{}
	b, ok := Bytes(f)
	if !ok {
		t.Fatal("Bytes reported false")
	}
	f.Kind = 0x01020304
	found := false
	for _, x := range b[:4] {
		if x != 0 {
			found = true
		}
	}
	if !found {
		t.Error("the view does not share the struct's memory")
	}
}

func TestBytesNil(t *testing.T) {
	if _, ok := Bytes(nil); ok {
		t.Error("Bytes(nil) reported ok, want false")
	}
}

func TestBytesDoesNotAllocate(t *testing.T) {
	f := &Frame{}
	var sink []byte
	if n := testing.AllocsPerRun(200, func() { sink, _ = Bytes(f) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
	_ = sink
}

func TestHasPointersFixture(t *testing.T) {
	type withPtr struct {
		A int
		B *int
	}
	type withString struct{ S string }
	type nested struct{ Inner withPtr }
	type clean struct {
		A int32
		B [4]byte
	}
	cases := []struct {
		v    any
		want bool
	}{
		{withPtr{}, true},
		{withString{}, true},
		{nested{}, true},
		{clean{}, false},
		{Frame{}, false},
	}
	for _, c := range cases {
		if got := hasPointers(reflect.TypeOf(c.v)); got != c.want {
			t.Errorf("hasPointers(%T) = %v, want %v", c.v, got, c.want)
		}
	}
}

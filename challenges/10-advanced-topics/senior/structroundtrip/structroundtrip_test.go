package structroundtrip

import (
	"testing"
	"unsafe"
)

func TestDecodeRoundTrip(t *testing.T) {
	in := Frame{Kind: 7, Seq: 9, Stamp: 1234567890}
	got, ok := Decode(Encode(&in))
	if !ok {
		t.Fatal("Decode reported false for a well-formed frame")
	}
	if got != in {
		t.Errorf("Decode = %+v, want %+v", got, in)
	}
}

func TestDecodeCopiesOut(t *testing.T) {
	in := Frame{Kind: 1}
	b := Encode(&in)
	got, ok := Decode(b)
	if !ok {
		t.Fatal("Decode reported false")
	}
	in.Kind = 99
	if got.Kind != 1 {
		t.Error("the result aliases the input bytes; it must be a copy")
	}
}

func TestDecodeWrongLength(t *testing.T) {
	in := Frame{}
	b := Encode(&in)
	for _, c := range [][]byte{nil, b[:4], b[:len(b)-1], append(append([]byte{}, b...), 0)} {
		if _, ok := Decode(c); ok {
			t.Errorf("Decode of %d bytes reported ok, want false", len(c))
		}
	}
}

func TestDecodeMisaligned(t *testing.T) {
	var zero Frame
	size := int(unsafe.Sizeof(zero))
	backing := make([]Frame, 2)
	all := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(backing))), size*2)
	if _, ok := Decode(all[1 : 1+size]); ok {
		t.Error("a misaligned frame reported ok, want false")
	}
}

func TestDecodeAllocatesNothing(t *testing.T) {
	in := Frame{Kind: 3}
	b := Encode(&in)
	var sink Frame
	if n := testing.AllocsPerRun(200, func() { sink, _ = Decode(b) }); n != 0 {
		t.Errorf("Decode made %v allocations, want 0", n)
	}
	_ = sink
}

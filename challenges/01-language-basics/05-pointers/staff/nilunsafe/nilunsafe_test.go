package nilunsafe

import (
	"testing"
	"unsafe"
)

func TestReadOr(t *testing.T) {
	if ReadOr(nil, 7) != 7 {
		t.Errorf("nil should give default (no panic)")
	}
	x := 5
	if ReadOr(unsafe.Pointer(&x), 7) != 5 {
		t.Errorf("want 5")
	}
}

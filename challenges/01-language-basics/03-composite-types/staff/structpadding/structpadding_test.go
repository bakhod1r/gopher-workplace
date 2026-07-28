package structpadding

import (
	"testing"
	"unsafe"
)

func TestSize(t *testing.T) {
	if got := unsafe.Sizeof(Record{}); got != 16 {
		t.Errorf("sizeof(Record)=%d; want 16 (reorder fields to reduce padding)", got)
	}
}

package alignof

import (
	"testing"
	"unsafe"
)

func TestFieldAlign(t *testing.T) {
	var s S
	if got := FieldAlign(&s); got != unsafe.Alignof(s.B) {
		t.Errorf("=%d want %d (use Alignof, not Sizeof)", got, unsafe.Alignof(s.B))
	}
}

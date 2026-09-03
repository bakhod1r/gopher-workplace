package upperascii

import (
	"bytes"
	"testing"
)

func TestUpper(t *testing.T) {
	if got := Upper([]byte("go1 x")); !bytes.Equal(got, []byte("GO1 X")) {
		t.Errorf("Upper = %q, want \"GO1 X\"", got)
	}
	if got := Upper([]byte("ALREADY")); !bytes.Equal(got, []byte("ALREADY")) {
		t.Errorf("Upper = %q, want \"ALREADY\"", got)
	}
	if got := Upper(nil); len(got) != 0 {
		t.Errorf("Upper(nil) = %q, want empty", got)
	}
}

func TestUpperIsInPlace(t *testing.T) {
	b := []byte("abc")
	out := Upper(b)
	if !bytes.Equal(b, []byte("ABC")) {
		t.Errorf("b = %q, want \"ABC\": the caller's buffer was not modified", b)
	}
	if &out[0] != &b[0] {
		t.Error("Upper returned a different array")
	}
}

func TestUpperAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("abcdef"), 32)
	if n := testing.AllocsPerRun(100, func() { _ = Upper(b) }); n != 0 {
		t.Errorf("Upper made %v allocations, want 0", n)
	}
}

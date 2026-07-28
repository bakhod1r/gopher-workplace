package strbytes

import "testing"

func TestByteLen(t *testing.T) {
	for _, s := range []string{"", "a", "hello", "héllo"} {
		if got := ByteLen(s); got != len(s) {
			t.Errorf("ByteLen(%q)=%d want %d", s, got, len(s))
		}
	}
}

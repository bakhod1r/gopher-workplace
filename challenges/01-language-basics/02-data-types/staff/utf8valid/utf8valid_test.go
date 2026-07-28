package utf8valid

import "testing"

func TestValid(t *testing.T) {
	cases := []struct {
		b    []byte
		want bool
	}{
		{[]byte("hello"), true},
		{[]byte("é"), true},         // C3 A9
		{[]byte("€"), true},         // E2 82 AC (3 bytes)
		{[]byte("日本語"), true},       // 3 bytes each
		{[]byte{0xC3}, false},       // truncated
		{[]byte{0xC3, 0x28}, false}, // bad continuation
		{[]byte{0x80}, false},       // lone continuation
	}
	for _, c := range cases {
		if got := Valid(c.b); got != c.want {
			t.Errorf("Valid(% x)=%v; want %v", c.b, got, c.want)
		}
	}
}

package perms

import "testing"

func TestBits(t *testing.T) {
	if Read != 1 || Write != 2 || Execute != 4 {
		t.Fatalf("bits = %d,%d,%d; want 1,2,4", Read, Write, Execute)
	}
}

func TestHas(t *testing.T) {
	rw := Read | Write
	cases := []struct {
		set, want Permission
		ok        bool
	}{
		{rw, Read, true},
		{rw, Write, true},
		{rw, Execute, false},
		{rw, Read | Write, true},
		{rw, Read | Execute, false},
		{0, Read, false},
	}
	for _, c := range cases {
		if got := Has(c.set, c.want); got != c.ok {
			t.Errorf("Has(%d,%d)=%v; want %v", c.set, c.want, got, c.ok)
		}
	}
}

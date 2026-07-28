package perms

import "testing"

func TestRevoke(t *testing.T) {
	all := Read | Write | Execute
	cases := []struct {
		set, drop, want Permission
	}{
		{all, Write, Read | Execute},
		{all, Read | Execute, Write},
		{Read, Write, Read}, // dropping absent bit is a no-op
		{all, all, 0},
	}
	for _, c := range cases {
		if got := Revoke(c.set, c.drop); got != c.want {
			t.Errorf("Revoke(%d,%d)=%d; want %d", c.set, c.drop, got, c.want)
		}
	}
}

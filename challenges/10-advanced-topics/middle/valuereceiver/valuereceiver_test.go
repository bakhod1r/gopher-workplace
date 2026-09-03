package valuereceiver

import "testing"

var sinkA, sinkB int

func TestTimeouts(t *testing.T) {
	c := &Config{Read: 5, Write: 9}
	if r, w := c.Timeouts(); r != 5 || w != 9 {
		t.Errorf("Timeouts = %d, %d, want 5, 9", r, w)
	}
	if r, w := (&Config{}).Timeouts(); r != 0 || w != 0 {
		t.Errorf("Timeouts = %d, %d, want 0, 0", r, w)
	}
}

func TestTimeoutsSeesLaterWrites(t *testing.T) {
	c := &Config{Read: 1}
	c.Read = 42
	if r, _ := c.Timeouts(); r != 42 {
		t.Errorf("read = %d, want 42: the receiver must be the caller's Config", r)
	}
}

func TestTimeoutsAllocatesNothing(t *testing.T) {
	c := &Config{Read: 1, Write: 2}
	if n := testing.AllocsPerRun(200, func() { sinkA, sinkB = c.Timeouts() }); n != 0 {
		t.Errorf("Timeouts made %v allocations, want 0", n)
	}
}

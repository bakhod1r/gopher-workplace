package ptrmethodsetbug

import "testing"

type counter struct {
	n int
}

func (c *counter) Reset() { c.n = 0 }

func TestResetAllMutatesSlice(t *testing.T) {
	s := []counter{{n: 3}, {n: 4}, {n: 5}}
	ResetAll[counter, *counter](s)
	for i, c := range s {
		if c.n != 0 {
			t.Errorf("s[%d].n = %d, want 0", i, c.n)
		}
	}
}

func TestResetAllSingle(t *testing.T) {
	s := []counter{{n: 9}}
	ResetAll[counter, *counter](s)
	if s[0].n != 0 {
		t.Errorf("s[0].n = %d, want 0", s[0].n)
	}
}

func TestResetAllEmpty(t *testing.T) {
	s := []counter{}
	ResetAll[counter, *counter](s)
	if len(s) != 0 {
		t.Errorf("len = %d, want 0", len(s))
	}
}

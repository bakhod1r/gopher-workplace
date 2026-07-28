package bareoverwrite

import "testing"

func TestDoubled(t *testing.T) {
	if got := Doubled(21); got != 42 {
		t.Errorf("=%d want 42", got)
	}
}

package methodcopy

import "testing"

func TestGetter(t *testing.T) {
	a := &Account{Balance: 100}
	get := Getter(a)
	a.Balance = 250
	if got := get(); got != 250 {
		t.Errorf("=%d want 250 (getter bound to a copy?)", got)
	}
}

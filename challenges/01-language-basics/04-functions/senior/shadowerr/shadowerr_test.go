package shadowerr

import "testing"

func TestParse(t *testing.T) {
	n, err := Parse("42")
	if n != 42 || err != nil {
		t.Errorf(`Parse("42")=%d,%v want 42,nil`, n, err)
	}
	_, err = Parse("nope")
	if err == nil {
		t.Errorf(`Parse("nope") should return a non-nil error`)
	}
}

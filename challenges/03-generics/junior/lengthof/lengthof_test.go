package lengthof

import "testing"

func TestLength(t *testing.T) {
	if got := Length("abc"); got != 3 {
		t.Errorf(`Length("abc") = %v, want 3`, got)
	}
	if got := Length([]byte{1, 2}); got != 2 {
		t.Errorf("Length([]byte{1, 2}) = %v, want 2", got)
	}
	if got := Length(""); got != 0 {
		t.Errorf(`Length("") = %v, want 0`, got)
	}
	if got := Length([]byte(nil)); got != 0 {
		t.Errorf("Length(nil bytes) = %v, want 0", got)
	}
}

package isempty

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty(nil) {
		t.Errorf("nil is empty")
	}
	if IsEmpty(&Node{}) {
		t.Errorf("non-nil is not empty")
	}
}

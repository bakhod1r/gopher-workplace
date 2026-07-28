package slicetoarray

import "testing"

func TestFirst4(t *testing.T) {
	a, ok := First4([]byte{1, 2, 3, 4, 5})
	if !ok || a != [4]byte{1, 2, 3, 4} {
		t.Errorf("got %v,%v", a, ok)
	}
	_, ok = First4([]byte{1, 2}) // too short: must be false, not a panic
	if ok {
		t.Error("short input must return false")
	}
}

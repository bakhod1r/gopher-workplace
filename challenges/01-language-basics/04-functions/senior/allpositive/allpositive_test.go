package allpositive

import "testing"

func TestAllPositive(t *testing.T) {
	if !AllPositive([]int{1, 2, 3}) {
		t.Errorf("all positive should be true")
	}
	if AllPositive([]int{1, -2, 3}) {
		t.Errorf("contains -2, should be false")
	}
	if !AllPositive(nil) {
		t.Errorf("empty should be true")
	}
}

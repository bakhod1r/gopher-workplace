package adder

import "testing"

func TestAdder(t *testing.T) {
	add10 := Adder(10)
	if add10(5) != 15 || add10(0) != 10 {
		t.Errorf("add10 wrong")
	}
	add1 := Adder(1)
	if add1(5) != 6 {
		t.Errorf("add1 wrong")
	}
}

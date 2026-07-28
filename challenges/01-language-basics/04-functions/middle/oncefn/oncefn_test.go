package oncefn

import "testing"

func TestOnce(t *testing.T) {
	n := 0
	do := Once(func() { n++ })
	do()
	do()
	do()
	if n != 1 {
		t.Errorf("f ran %d times, want 1", n)
	}
}

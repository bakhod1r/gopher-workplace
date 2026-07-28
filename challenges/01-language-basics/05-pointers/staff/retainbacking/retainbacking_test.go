package retainbacking

import "testing"

func TestPrefix(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	p := Prefix(xs, 2)
	_ = append(p, 99) // must not overwrite xs[2]
	if xs[2] != 3 {
		t.Errorf("xs[2]=%d want 3 (append spilled into parent)", xs[2])
	}
}

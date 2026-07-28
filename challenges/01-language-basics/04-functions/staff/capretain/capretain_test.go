package capretain

import "testing"

func TestHead(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	h := Head(xs, 2)
	_ = append(h, 99) // must NOT overwrite xs[2]
	if xs[2] != 3 {
		t.Errorf("xs[2]=%d want 3 (append into shared spare cap clobbered it)", xs[2])
	}
}

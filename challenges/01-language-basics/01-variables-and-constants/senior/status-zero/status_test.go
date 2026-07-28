package status

import "testing"

func TestZeroUnknown(t *testing.T) {
	var s Status // zero value
	if IsKnown(s) {
		t.Fatal("zero-valued Status must be unknown")
	}
}

func TestKnown(t *testing.T) {
	for _, s := range []Status{Pending, Shipped, Delivered} {
		if !IsKnown(s) {
			t.Errorf("%d should be known", s)
		}
	}
	if Pending != 1 || Shipped != 2 || Delivered != 3 {
		t.Fatalf("values=%d,%d,%d; want 1,2,3", Pending, Shipped, Delivered)
	}
}

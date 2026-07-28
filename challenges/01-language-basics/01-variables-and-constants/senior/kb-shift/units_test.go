package units

import "testing"

func TestScale(t *testing.T) {
	if KB != 1024 {
		t.Fatalf("KB=%d; want 1024", KB)
	}
	if MB != 1024*1024 {
		t.Fatalf("MB=%d; want %d", MB, 1024*1024)
	}
	if GB != 1024*1024*1024 {
		t.Fatalf("GB=%d; want %d", GB, 1024*1024*1024)
	}
}

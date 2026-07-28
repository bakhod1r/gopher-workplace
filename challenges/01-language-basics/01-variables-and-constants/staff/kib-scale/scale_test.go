package scale

import "testing"

func TestKiB(t *testing.T) {
	if KiB != 1024 {
		t.Fatalf("KiB=%d; want 1024", KiB)
	}
	if Bytes(2) != 2048 {
		t.Fatalf("Bytes(2)=%d; want 2048", Bytes(2))
	}
}

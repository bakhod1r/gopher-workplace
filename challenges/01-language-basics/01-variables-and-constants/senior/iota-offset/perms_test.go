package perms

import "testing"

func TestBits(t *testing.T) {
	if Read != 1 || Write != 2 || Execute != 4 {
		t.Fatalf("bits=%d,%d,%d; want 1,2,4", Read, Write, Execute)
	}
}

func TestHas(t *testing.T) {
	if !Has(Read|Write, Read) || Has(Read|Write, Execute) {
		t.Error("membership wrong")
	}
}

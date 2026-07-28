package stringbyte

import "testing"

func TestFirstByte(t *testing.T) {
	if got := FirstByte("Xyz"); got != 'X' {
		t.Errorf("=%q want 'X' (read the header, not the data?)", got)
	}
}

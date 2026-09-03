package ifaceescape

import "testing"

var sinkN int

func TestChecksum(t *testing.T) {
	if got := Checksum([]int{1}); got != int('1') {
		t.Errorf("Checksum = %d, want %d", got, int('1'))
	}
	if got := Checksum(nil); got != 0 {
		t.Errorf("Checksum = %d, want 0", got)
	}
	if got := Checksum([]int{12}); got != int('1')+int('2') {
		t.Errorf("Checksum = %d, want %d", got, int('1')+int('2'))
	}
	if got := Checksum([]int{-1}); got != int('-')+int('1') {
		t.Errorf("Checksum = %d, want %d", got, int('-')+int('1'))
	}
}

func TestChecksumAllocatesNothing(t *testing.T) {
	vals := []int{1, 22, 333, 4444}
	if n := testing.AllocsPerRun(200, func() { sinkN = Checksum(vals) }); n != 0 {
		t.Errorf("Checksum made %v allocations, want 0: the scratch buffer is escaping", n)
	}
}

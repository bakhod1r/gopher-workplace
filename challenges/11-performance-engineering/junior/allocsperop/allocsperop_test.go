package allocsperop

import "testing"

var sink []byte
var sum int

func TestAllocsOfCountsHeapAllocations(t *testing.T) {
	got := AllocsOf(100, func() { sink = make([]byte, 1024) })
	if got != 1 {
		t.Errorf("AllocsOf(make) = %d, want 1", got)
	}
}

func TestAllocsOfZeroForAllocationFreeWork(t *testing.T) {
	got := AllocsOf(100, func() { sum += 1 })
	if got != 0 {
		t.Errorf("AllocsOf(add) = %d, want 0", got)
	}
}

func TestAllocsOfCountsEachAllocation(t *testing.T) {
	got := AllocsOf(100, func() {
		sink = make([]byte, 1024)
		sink = make([]byte, 2048)
		sink = make([]byte, 4096)
	})
	if got != 3 {
		t.Errorf("AllocsOf(3 makes) = %d, want 3", got)
	}
}

func TestAllocsOfNonPositiveRuns(t *testing.T) {
	if got := AllocsOf(0, func() { sink = make([]byte, 1024) }); got != 1 {
		t.Errorf("AllocsOf(0, make) = %d, want 1", got)
	}
}

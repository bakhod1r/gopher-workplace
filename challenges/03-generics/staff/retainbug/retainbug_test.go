package retainbug

import (
	"reflect"
	"runtime"
	"testing"
)

func TestHeadContents(t *testing.T) {
	if got := Head([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Head = %v, want [1 2]", got)
	}
	if got := Head([]int{1, 2}, 9); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Head = %v, want [1 2]", got)
	}
	if got := Head([]int{1, 2}, -1); len(got) != 0 {
		t.Errorf("Head = %v, want []", got)
	}
}

func TestHeadIsIndependentOfInput(t *testing.T) {
	s := []int{1, 2, 3, 4}
	h := Head(s, 2)
	s[0] = 99
	if h[0] != 1 {
		t.Errorf("h[0] = %d, want 1: the head still views the input's storage", h[0])
	}
}

func TestHeadReleasesTheBackingArray(t *testing.T) {
	const chunks = 60
	const size = 1 << 20 // 8 MB of int per payload
	const ceiling = 100 << 20

	kept := make([][]int, 0, chunks)
	for i := 0; i < chunks; i++ {
		payload := make([]int, size)
		payload[0] = i
		kept = append(kept, Head(payload, 4))
	}

	runtime.GC()
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	if kept[7][0] != 7 {
		t.Fatalf("kept[7][0] = %d, want 7", kept[7][0])
	}
	if ms.HeapAlloc > ceiling {
		t.Errorf("heap holds %d MB after dropping %d payloads, ceiling %d MB: the heads retain their payloads",
			ms.HeapAlloc>>20, chunks, ceiling>>20)
	}
}

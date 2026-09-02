package growbug

import (
	"reflect"
	"testing"
)

func TestCollectContent(t *testing.T) {
	if got := Collect([]int{1}, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Collect = %v, want [1 2 3]", got)
	}
	if got := Collect([]int(nil), 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Collect(nil, 1) = %v, want [1]", got)
	}
}

func TestCollectReservesOnce(t *testing.T) {
	base := make([]int, 0, 1)
	base = append(base, 0)
	got := Collect(base, 1, 2, 3, 4, 5, 6, 7)
	if cap(got) < len(got) {
		t.Fatalf("cap %d < len %d", cap(got), len(got))
	}
	if len(got) != 8 {
		t.Fatalf("len = %d, want 8", len(got))
	}
	for i, want := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
		if got[i] != want {
			t.Fatalf("got[%d] = %v, want %v", i, got[i], want)
		}
	}
}

package unsafestride

import "testing"

func TestAt(t *testing.T) {
	arr := [4]int32{10, 20, 30, 40}
	for i := 0; i < 4; i++ {
		if got := At(&arr, i); got != arr[i] {
			t.Errorf("At(%d)=%d want %d", i, got, arr[i])
		}
	}
}

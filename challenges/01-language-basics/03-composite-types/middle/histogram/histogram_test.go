package histogram

import (
	"reflect"
	"testing"
)

func TestBucket(t *testing.T) {
	// size 10: bins [0,10),[10,20),[20,30)
	got := Bucket([]int{5, 12, 15, 25, 3}, 10)
	want := []int{2, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bucket=%v; want %v", got, want)
	}
	if len(Bucket(nil, 10)) != 0 {
		t.Error("empty -> empty")
	}
}

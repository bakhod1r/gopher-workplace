package wordwrap

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	// width 10
	got := Wrap("the quick brown fox", 10)
	want := []string{"the quick", "brown fox"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrap=%v; want %v", got, want)
	}
	// "aa bb cc" width 5 -> "aa bb","cc"
	got = Wrap("aa bb cc", 5)
	want = []string{"aa bb", "cc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrap=%v; want %v", got, want)
	}
}

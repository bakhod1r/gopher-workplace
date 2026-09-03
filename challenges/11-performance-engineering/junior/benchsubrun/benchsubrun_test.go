package benchsubrun

import (
	"reflect"
	"testing"
)

func TestNames(t *testing.T) {
	got := Names("BenchmarkEncode", []int{1, 10, 100})
	want := []string{"BenchmarkEncode/size=1", "BenchmarkEncode/size=10", "BenchmarkEncode/size=100"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Names = %v, want %v", got, want)
	}
}

func TestNamesKeepsOrder(t *testing.T) {
	got := Names("B", []int{100, 1})
	want := []string{"B/size=100", "B/size=1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Names = %v, want %v", got, want)
	}
}

func TestNamesEmpty(t *testing.T) {
	got := Names("B", nil)
	if got == nil {
		t.Fatal("Names(nil sizes) = nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("Names = %v, want empty", got)
	}
}

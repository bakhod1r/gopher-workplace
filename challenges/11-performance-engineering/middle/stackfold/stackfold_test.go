package stackfold

import (
	"reflect"
	"testing"
)

func TestFoldSumsIdenticalStacks(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"a", "b"}, 3},
		{[]string{"a", "b"}, 2},
	})
	if !reflect.DeepEqual(got, []string{"a;b 5"}) {
		t.Errorf("Fold = %v, want [a;b 5]", got)
	}
}

func TestFoldOrdersByValueThenStack(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"z"}, 1},
		{[]string{"a", "b"}, 9},
		{[]string{"m"}, 1},
	})
	want := []string{"a;b 9", "m 1", "z 1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fold = %v, want %v", got, want)
	}
}

func TestFoldDistinguishesDifferentStacks(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"a", "b"}, 1},
		{[]string{"b", "a"}, 2},
	})
	want := []string{"b;a 2", "a;b 1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fold = %v, want %v", got, want)
	}
}

func TestFoldDropsJunk(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"a"}, 0},
		{[]string{"b"}, -1},
		{nil, 5},
	})
	if got == nil || len(got) != 0 {
		t.Errorf("Fold = %v, want empty non-nil slice", got)
	}
}

func TestFoldIsDeterministic(t *testing.T) {
	samples := []Sample{
		{[]string{"a"}, 2}, {[]string{"b"}, 2}, {[]string{"c"}, 2}, {[]string{"d"}, 2},
	}
	first := Fold(samples)
	for i := 0; i < 20; i++ {
		if !reflect.DeepEqual(Fold(samples), first) {
			t.Fatal("Fold is not deterministic")
		}
	}
}

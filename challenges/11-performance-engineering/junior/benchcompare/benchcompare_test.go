package benchcompare

import (
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	got := Compare(
		map[string]float64{"A": 100, "B": 200},
		map[string]float64{"A": 80, "B": 250},
	)
	want := []Delta{
		{"A", 100, 80, -20},
		{"B", 200, 250, 25},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Compare = %v, want %v", got, want)
	}
}

func TestCompareOnlyPairedNames(t *testing.T) {
	got := Compare(
		map[string]float64{"A": 100, "OnlyBase": 5},
		map[string]float64{"A": 100, "OnlyNew": 5},
	)
	want := []Delta{{"A", 100, 100, 0}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Compare = %v, want %v", got, want)
	}
}

func TestCompareIsSortedAndDeterministic(t *testing.T) {
	base := map[string]float64{"z": 10, "a": 10, "m": 10}
	cand := map[string]float64{"z": 5, "a": 5, "m": 5}
	first := Compare(base, cand)
	if len(first) != 3 || first[0].Name != "a" || first[1].Name != "m" || first[2].Name != "z" {
		t.Fatalf("Compare = %v, want rows ordered a, m, z", first)
	}
	for i := 0; i < 20; i++ {
		if !reflect.DeepEqual(Compare(base, cand), first) {
			t.Fatal("Compare is not deterministic across map iteration orders")
		}
	}
}

func TestCompareSkipsUnusableBaselines(t *testing.T) {
	got := Compare(
		map[string]float64{"A": 0, "B": -1, "C": 100},
		map[string]float64{"A": 5, "B": 5, "C": 50},
	)
	want := []Delta{{"C", 100, 50, -50}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Compare = %v, want %v", got, want)
	}
}

func TestCompareEmpty(t *testing.T) {
	got := Compare(nil, nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Compare(nil, nil) = %v, want empty non-nil slice", got)
	}
}

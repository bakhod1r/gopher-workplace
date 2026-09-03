package profilelabelsum

import (
	"reflect"
	"testing"
)

func TestSumByLabel(t *testing.T) {
	got := SumByLabel([]Sample{
		{map[string]string{"kind": "read"}, 3},
		{map[string]string{"kind": "write"}, 4},
		{map[string]string{"kind": "read"}, 2},
	}, "kind")
	want := map[string]int64{"read": 5, "write": 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumByLabel = %v, want %v", got, want)
	}
}

func TestSumByLabelMissingKeyGroupsUnderEmpty(t *testing.T) {
	got := SumByLabel([]Sample{
		{map[string]string{"kind": "read"}, 3},
		{map[string]string{"tenant": "acme"}, 4},
		{nil, 1},
	}, "kind")
	want := map[string]int64{"read": 3, "": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumByLabel = %v, want %v", got, want)
	}
}

func TestSumByLabelIgnoresOtherKeys(t *testing.T) {
	got := SumByLabel([]Sample{
		{map[string]string{"kind": "read", "tenant": "acme"}, 3},
	}, "tenant")
	want := map[string]int64{"acme": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumByLabel = %v, want %v", got, want)
	}
}

func TestSumByLabelSkipsNonPositive(t *testing.T) {
	got := SumByLabel([]Sample{
		{map[string]string{"kind": "read"}, 0},
		{map[string]string{"kind": "read"}, -2},
	}, "kind")
	if got == nil || len(got) != 0 {
		t.Errorf("SumByLabel = %v, want empty non-nil map", got)
	}
}

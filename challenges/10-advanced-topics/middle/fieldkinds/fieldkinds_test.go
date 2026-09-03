package fieldkinds

import (
	"reflect"
	"testing"
)

type row struct {
	ID     int
	Name   string
	Tags   []string
	Ratio  float64
	hidden bool
}

func TestFieldKinds(t *testing.T) {
	want := []string{"ID:int", "Name:string", "Tags:slice", "Ratio:float64"}
	if got := FieldKinds(row{}); !reflect.DeepEqual(got, want) {
		t.Errorf("FieldKinds = %v, want %v", got, want)
	}
}

func TestFieldKindsNested(t *testing.T) {
	type outer struct {
		In  row
		Ptr *row
	}
	want := []string{"In:struct", "Ptr:ptr"}
	if got := FieldKinds(outer{}); !reflect.DeepEqual(got, want) {
		t.Errorf("FieldKinds = %v, want %v", got, want)
	}
}

func TestFieldKindsRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, map[string]int{}, &row{}} {
		if got := FieldKinds(in); got != nil {
			t.Errorf("FieldKinds(%#v) = %v, want nil", in, got)
		}
	}
}

func TestFieldKindsEmptyStruct(t *testing.T) {
	if got := FieldKinds(struct{}{}); got != nil {
		t.Errorf("FieldKinds = %v, want nil", got)
	}
}

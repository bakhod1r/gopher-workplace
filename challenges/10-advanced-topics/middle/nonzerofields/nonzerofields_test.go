package nonzerofields

import (
	"reflect"
	"testing"
)

type patch struct {
	Name   string
	Count  int
	Active bool
	Tags   []string
	hidden int
}

func TestNonZero(t *testing.T) {
	got := NonZero(patch{Name: "x", Count: 3})
	if !reflect.DeepEqual(got, []string{"Name", "Count"}) {
		t.Errorf("NonZero = %v, want [Name Count]", got)
	}
}

func TestNonZeroAllZero(t *testing.T) {
	if got := NonZero(patch{}); got != nil {
		t.Errorf("NonZero = %v, want nil", got)
	}
}

func TestNonZeroSkipsUnexported(t *testing.T) {
	if got := NonZero(patch{hidden: 9}); got != nil {
		t.Errorf("NonZero = %v, want nil: unexported fields do not count", got)
	}
}

func TestNonZeroEmptySliceCounts(t *testing.T) {
	got := NonZero(patch{Tags: []string{}})
	if !reflect.DeepEqual(got, []string{"Tags"}) {
		t.Errorf("NonZero = %v, want [Tags]: an empty non-nil slice is not the zero value", got)
	}
}

func TestNonZeroFalseIsZero(t *testing.T) {
	if got := NonZero(patch{Active: false}); got != nil {
		t.Errorf("NonZero = %v, want nil", got)
	}
	if got := NonZero(patch{Active: true}); !reflect.DeepEqual(got, []string{"Active"}) {
		t.Errorf("NonZero = %v, want [Active]", got)
	}
}

func TestNonZeroNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &patch{}} {
		if got := NonZero(in); got != nil {
			t.Errorf("NonZero(%#v) = %v, want nil", in, got)
		}
	}
}

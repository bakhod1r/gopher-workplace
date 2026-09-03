package fieldnames

import (
	"reflect"
	"testing"
)

type user struct {
	Name  string
	Age   int
	admin bool
}

func TestFieldNames(t *testing.T) {
	if got := FieldNames(user{}); !reflect.DeepEqual(got, []string{"Name", "Age"}) {
		t.Errorf("FieldNames = %v, want [Name Age]", got)
	}
	if got := FieldNames(struct{}{}); got != nil {
		t.Errorf("FieldNames = %v, want nil for an empty struct", got)
	}
}

func TestFieldNamesRejectsNonStructs(t *testing.T) {
	for _, in := range []any{3, "s", []int{1}, map[string]int{}, nil, &user{}} {
		if got := FieldNames(in); got != nil {
			t.Errorf("FieldNames(%#v) = %v, want nil", in, got)
		}
	}
}

func TestFieldNamesIsInDeclarationOrder(t *testing.T) {
	type ordered struct {
		Z, A, M int
	}
	if got := FieldNames(ordered{}); !reflect.DeepEqual(got, []string{"Z", "A", "M"}) {
		t.Errorf("FieldNames = %v, want [Z A M]", got)
	}
}

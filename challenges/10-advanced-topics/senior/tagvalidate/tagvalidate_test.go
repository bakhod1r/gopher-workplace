package tagvalidate

import (
	"reflect"
	"testing"
)

type good struct {
	A      int    `col:"a"`
	B      string `col:"b"`
	hidden int
}

type missing struct {
	A int `col:"a"`
	B int
}

type empty struct {
	A int `col:""`
}

type dup struct {
	A int `col:"x"`
	B int `col:"x"`
}

type comma struct {
	A int `col:"a,omitempty"`
}

func TestValidateGood(t *testing.T) {
	if got := Validate(good{}); len(got) != 0 {
		t.Errorf("Validate = %v, want no problems", got)
	}
}

func TestValidateMissing(t *testing.T) {
	if got := Validate(missing{}); !reflect.DeepEqual(got, []string{"B: missing col tag"}) {
		t.Errorf("Validate = %v, want [B: missing col tag]", got)
	}
	if got := Validate(empty{}); !reflect.DeepEqual(got, []string{"A: missing col tag"}) {
		t.Errorf("Validate = %v, want [A: missing col tag]", got)
	}
}

func TestValidateDuplicate(t *testing.T) {
	if got := Validate(dup{}); !reflect.DeepEqual(got, []string{"B: duplicate tag of A"}) {
		t.Errorf("Validate = %v, want [B: duplicate tag of A]", got)
	}
}

func TestValidateComma(t *testing.T) {
	if got := Validate(comma{}); !reflect.DeepEqual(got, []string{"A: tag contains a comma"}) {
		t.Errorf("Validate = %v, want [A: tag contains a comma]", got)
	}
}

func TestValidateNonStruct(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}} {
		if got := Validate(in); len(got) != 1 || got[0] != "not a struct" {
			t.Errorf("Validate(%#v) = %v, want [not a struct]", in, got)
		}
	}
}

func TestValidateReportsInFieldOrder(t *testing.T) {
	type multi struct {
		A int `col:"x"`
		B int
		C int `col:"x"`
	}
	want := []string{"B: missing col tag", "C: duplicate tag of A"}
	if got := Validate(multi{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Validate = %v, want %v", got, want)
	}
}

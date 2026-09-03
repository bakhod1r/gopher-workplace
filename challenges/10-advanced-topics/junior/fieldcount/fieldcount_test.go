package fieldcount

import "testing"

type rec struct {
	A      int
	B      string
	hidden bool
}

func TestFieldCount(t *testing.T) {
	total, exported := FieldCount(rec{})
	if total != 3 || exported != 2 {
		t.Errorf("FieldCount = %d, %d, want 3, 2", total, exported)
	}
}

func TestFieldCountEmptyStruct(t *testing.T) {
	total, exported := FieldCount(struct{}{})
	if total != 0 || exported != 0 {
		t.Errorf("FieldCount = %d, %d, want 0, 0", total, exported)
	}
}

func TestFieldCountAllUnexported(t *testing.T) {
	type private struct{ a, b int }
	total, exported := FieldCount(private{})
	if total != 2 || exported != 0 {
		t.Errorf("FieldCount = %d, %d, want 2, 0", total, exported)
	}
}

func TestFieldCountNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, "s", []int{1}, &rec{}} {
		total, exported := FieldCount(in)
		if total != 0 || exported != 0 {
			t.Errorf("FieldCount(%#v) = %d, %d, want 0, 0", in, total, exported)
		}
	}
}

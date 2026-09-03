package zerostrings

import (
	"errors"
	"testing"
)

type inner struct {
	Secret string
	Count  int
}

type record struct {
	Name   string
	In     inner
	Ptr    *inner
	List   []inner
	hidden string
}

func TestRedactFlat(t *testing.T) {
	r := &record{Name: "top", hidden: "keep"}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.Name != "" {
		t.Errorf("Name = %q, want empty", r.Name)
	}
	if r.hidden != "keep" {
		t.Errorf("hidden = %q, want \"keep\"", r.hidden)
	}
}

func TestRedactNested(t *testing.T) {
	r := &record{
		Name: "top",
		In:   inner{Secret: "a", Count: 1},
		Ptr:  &inner{Secret: "b"},
		List: []inner{{Secret: "c"}, {Secret: "d"}},
	}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.In.Secret != "" || r.Ptr.Secret != "" {
		t.Errorf("nested secrets survived: %+v", r)
	}
	for i, e := range r.List {
		if e.Secret != "" {
			t.Errorf("List[%d].Secret = %q, want empty", i, e.Secret)
		}
	}
	if r.In.Count != 1 {
		t.Errorf("Count = %d, want 1: non-string fields must be untouched", r.In.Count)
	}
}

func TestRedactNilPointerField(t *testing.T) {
	r := &record{Name: "x"}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.Ptr != nil {
		t.Error("the nil pointer field was replaced")
	}
}

func TestRedactBadTarget(t *testing.T) {
	for _, c := range []any{record{}, nil, (*record)(nil), new(int)} {
		if err := Redact(c); !errors.Is(err, ErrTarget) {
			t.Errorf("Redact(%#v) = %v, want ErrTarget", c, err)
		}
	}
}

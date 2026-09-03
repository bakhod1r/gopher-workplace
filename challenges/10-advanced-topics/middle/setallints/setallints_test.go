package setallints

import (
	"errors"
	"testing"
)

type rec struct {
	A      int
	B      int
	Name   string
	hidden int
	Ratio  float64
}

func TestSetAllInts(t *testing.T) {
	r := &rec{Name: "keep", hidden: 1}
	n, err := SetAllInts(r, 7)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Errorf("count = %d, want 2", n)
	}
	if r.A != 7 || r.B != 7 {
		t.Errorf("r = %+v, want A and B set to 7", *r)
	}
	if r.Name != "keep" || r.hidden != 1 || r.Ratio != 0 {
		t.Errorf("r = %+v: other fields must be untouched", *r)
	}
}

func TestSetAllIntsNoIntFields(t *testing.T) {
	var s struct {
		A string
		B bool
	}
	n, err := SetAllInts(&s, 1)
	if err != nil || n != 0 {
		t.Errorf("count = %d, err = %v, want 0, nil", n, err)
	}
}

func TestSetAllIntsBadTarget(t *testing.T) {
	for _, c := range []any{rec{}, nil, (*rec)(nil), new(int)} {
		if _, err := SetAllInts(c, 1); !errors.Is(err, ErrTarget) {
			t.Errorf("SetAllInts(%#v) = %v, want ErrTarget", c, err)
		}
	}
}

func TestSetAllIntsOverwrites(t *testing.T) {
	r := &rec{A: 1, B: 2}
	if _, err := SetAllInts(r, 0); err != nil {
		t.Fatal(err)
	}
	if r.A != 0 || r.B != 0 {
		t.Errorf("r = %+v, want both zeroed", *r)
	}
}

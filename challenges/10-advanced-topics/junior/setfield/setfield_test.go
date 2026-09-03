package setfield

import (
	"errors"
	"testing"
)

type counters struct {
	Hits   int
	Name   string
	hidden int
}

func TestSetInt(t *testing.T) {
	c := &counters{}
	if err := SetInt(c, "Hits", 42); err != nil {
		t.Fatalf("SetInt = %v, want nil", err)
	}
	if c.Hits != 42 {
		t.Errorf("Hits = %d, want 42", c.Hits)
	}
}

func TestSetIntRejectsBadTargets(t *testing.T) {
	cases := []struct {
		name string
		ptr  any
		f    string
	}{
		{"value not pointer", counters{}, "Hits"},
		{"nil pointer", (*counters)(nil), "Hits"},
		{"nil interface", nil, "Hits"},
		{"missing field", &counters{}, "Nope"},
		{"wrong kind", &counters{}, "Name"},
		{"unexported", &counters{}, "hidden"},
		{"not a struct", new(int), "Hits"},
	}
	for _, c := range cases {
		if err := SetInt(c.ptr, c.f, 1); !errors.Is(err, ErrNotSettable) {
			t.Errorf("%s: err = %v, want ErrNotSettable", c.name, err)
		}
	}
}

func TestSetIntOverwrites(t *testing.T) {
	c := &counters{Hits: 5}
	if err := SetInt(c, "Hits", 0); err != nil || c.Hits != 0 {
		t.Errorf("Hits = %d, err = %v, want 0, nil", c.Hits, err)
	}
}

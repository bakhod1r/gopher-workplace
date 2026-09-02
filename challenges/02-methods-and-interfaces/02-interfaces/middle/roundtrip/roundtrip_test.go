package roundtrip

import (
	"errors"
	"testing"
)

func TestMarshal(t *testing.T) {
	if got := (&Record{ID: 1, Name: "a"}).Marshal(); got != "1|a" {
		t.Errorf("Marshal = %q, want \"1|a\"", got)
	}
	if got := (&Record{}).Marshal(); got != "0|" {
		t.Errorf("Marshal = %q, want \"0|\"", got)
	}
}

func TestUnmarshal(t *testing.T) {
	var r Record
	if err := r.Unmarshal("2|b"); err != nil {
		t.Fatalf("err = %v", err)
	}
	if r.ID != 2 || r.Name != "b" {
		t.Errorf("got %+v, want {2 b}", r)
	}

	if err := r.Unmarshal("oops"); !errors.Is(err, ErrBadFormat) {
		t.Errorf("err = %v, want ErrBadFormat", err)
	}
	if err := r.Unmarshal("x|b"); !errors.Is(err, ErrBadFormat) {
		t.Errorf("err = %v, want ErrBadFormat", err)
	}
	if err := r.Unmarshal("3|a|b"); err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if r.Name != "a|b" {
		t.Errorf("Name = %q, want \"a|b\"", r.Name)
	}
}

func TestRoundTrip(t *testing.T) {
	ok, err := RoundTrip(&Record{ID: 3, Name: "c"}, &Record{})
	if err != nil || !ok {
		t.Errorf("RoundTrip = %v, %v; want true, nil", ok, err)
	}

	ok, err = RoundTrip(&Record{ID: 0, Name: ""}, &Record{})
	if err != nil || !ok {
		t.Errorf("RoundTrip zero = %v, %v; want true, nil", ok, err)
	}
}

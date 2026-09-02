package unmarshaler

import (
	"errors"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var p Pair
	if err := p.Unmarshal("a=1"); err != nil {
		t.Fatalf("Unmarshal = %v, want nil", err)
	}
	if p.Key != "a" || p.Value != "1" {
		t.Errorf("got %+v, want {a 1}", p)
	}

	if err := p.Unmarshal("nope"); !errors.Is(err, ErrBadPair) {
		t.Errorf("Unmarshal = %v, want ErrBadPair", err)
	}

	if err := p.Unmarshal("k=v=w"); err != nil {
		t.Fatalf("Unmarshal = %v, want nil", err)
	}
	if p.Value != "v=w" {
		t.Errorf("Value = %q, want \"v=w\"", p.Value)
	}

	if err := p.Unmarshal("k="); err != nil {
		t.Fatalf("Unmarshal = %v, want nil", err)
	}
	if p.Value != "" {
		t.Errorf("Value = %q, want empty", p.Value)
	}
}

func TestUnmarshalAll(t *testing.T) {
	got, err := UnmarshalAll([]string{"a=1", "b=2"})
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if len(got) != 2 || got[1].Key != "b" {
		t.Errorf("got %+v", got)
	}

	if _, err := UnmarshalAll([]string{"a=1", "bad"}); !errors.Is(err, ErrBadPair) {
		t.Errorf("err = %v, want ErrBadPair", err)
	}

	got, err = UnmarshalAll(nil)
	if err != nil || len(got) != 0 {
		t.Errorf("UnmarshalAll(nil) = %v, %v", got, err)
	}
}

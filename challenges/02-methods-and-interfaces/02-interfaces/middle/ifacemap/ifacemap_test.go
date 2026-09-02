package ifacemap

import (
	"errors"
	"strings"
	"testing"
)

func TestRunRegistered(t *testing.T) {
	r := NewRegistry()
	r.Register("up", CommandFunc(strings.ToUpper))

	got, err := r.Run("up", "hi")
	if err != nil || got != "HI" {
		t.Errorf("Run = %q, %v; want \"HI\", nil", got, err)
	}
}

func TestRunUnknown(t *testing.T) {
	r := NewRegistry()
	got, err := r.Run("nope", "hi")
	if !errors.Is(err, ErrUnknown) {
		t.Errorf("err = %v, want ErrUnknown", err)
	}
	if got != "" {
		t.Errorf("got = %q, want empty", got)
	}
}

func TestRegisterReplaces(t *testing.T) {
	r := NewRegistry()
	r.Register("x", CommandFunc(strings.ToUpper))
	r.Register("x", CommandFunc(strings.ToLower))

	if got, _ := r.Run("x", "AB"); got != "ab" {
		t.Errorf("Run = %q, want \"ab\"", got)
	}
	if n := len(r.Names()); n != 1 {
		t.Errorf("Names len = %d, want 1", n)
	}
}

func TestNamesSorted(t *testing.T) {
	r := NewRegistry()
	r.Register("b", CommandFunc(strings.ToUpper))
	r.Register("a", CommandFunc(strings.ToUpper))
	r.Register("c", CommandFunc(strings.ToUpper))

	got := r.Names()
	want := []string{"a", "b", "c"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Names = %v, want %v", got, want)
		}
	}

	if n := len(NewRegistry().Names()); n != 0 {
		t.Errorf("empty Names len = %d, want 0", n)
	}
}

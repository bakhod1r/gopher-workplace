package grouprules

import (
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {
	errShort := errors.New("too short")
	errDigit := errors.New("needs a digit")

	pass := Rule{Name: "ok", Fn: func(string) error { return nil }}
	short := Rule{Name: "len", Fn: func(string) error { return errShort }}
	digit := Rule{Name: "digit", Fn: func(string) error { return errDigit }}

	t.Run("no_rules", func(t *testing.T) {
		if err := Check("abc", nil); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("all_pass", func(t *testing.T) {
		if err := Check("abc", []Rule{pass, pass}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("names_the_rule", func(t *testing.T) {
		err := Check("a", []Rule{short})
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if err.Error() != "len: too short" {
			t.Errorf("message = %q, want %q", err.Error(), "len: too short")
		}
	})

	t.Run("runs_every_rule", func(t *testing.T) {
		ran := 0
		count := Rule{Name: "count", Fn: func(string) error { ran++; return nil }}
		err := Check("a", []Rule{short, count, digit})
		if ran != 1 {
			t.Fatalf("later rule did not run")
		}
		if !errors.Is(err, errShort) || !errors.Is(err, errDigit) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
		if err.Error() != "len: too short\ndigit: needs a digit" {
			t.Errorf("message = %q, want the failures in declaration order", err.Error())
		}
	})
}

package mustatoi

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestParseInt(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		got, err := ParseInt("42")
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if got != 42 {
			t.Errorf("got = %d, want 42", got)
		}
	})

	t.Run("negative", func(t *testing.T) {
		got, err := ParseInt("-7")
		if err != nil || got != -7 {
			t.Errorf("ParseInt(\"-7\") = %d, %v, want -7, nil", got, err)
		}
	})

	t.Run("invalid_value", func(t *testing.T) {
		got, err := ParseInt("x")
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if got != 0 {
			t.Errorf("got = %d, want 0", got)
		}
	})

	t.Run("invalid_message", func(t *testing.T) {
		_, err := ParseInt("x")
		if !strings.HasPrefix(err.Error(), `parse "x": `) {
			t.Errorf("message = %q, want prefix %q", err.Error(), `parse "x": `)
		}
	})

	t.Run("keeps_syntax_sentinel", func(t *testing.T) {
		_, err := ParseInt("x")
		if !errors.Is(err, strconv.ErrSyntax) {
			t.Errorf("errors.Is(%v, strconv.ErrSyntax) = false, want true", err)
		}
	})
}

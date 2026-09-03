package parseerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseError(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		err := &ParseError{Line: 4, Msg: "bad token"}
		if err.Error() != "line 4: bad token" {
			t.Errorf("Error() = %q, want %q", err.Error(), "line 4: bad token")
		}
	})

	t.Run("direct", func(t *testing.T) {
		line, ok := LineOf(&ParseError{Line: 4, Msg: "bad token"})
		if !ok || line != 4 {
			t.Errorf("LineOf = %d, %v, want 4, true", line, ok)
		}
	})

	t.Run("wrapped", func(t *testing.T) {
		wrapped := fmt.Errorf("config: %w", &ParseError{Line: 9, Msg: "eof"})
		line, ok := LineOf(wrapped)
		if !ok || line != 9 {
			t.Errorf("LineOf = %d, %v, want 9, true", line, ok)
		}
	})

	t.Run("absent", func(t *testing.T) {
		line, ok := LineOf(errors.New("boom"))
		if ok || line != 0 {
			t.Errorf("LineOf = %d, %v, want 0, false", line, ok)
		}
	})

	t.Run("nil", func(t *testing.T) {
		if line, ok := LineOf(nil); ok || line != 0 {
			t.Errorf("LineOf(nil) = %d, %v, want 0, false", line, ok)
		}
	})
}

package stackframes

import (
	"strings"
	"testing"
)

func TestFrames(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		if got := Frames(0); len(got) != 0 {
			t.Errorf("Frames(0) = %v, want empty", got)
		}
	})

	t.Run("negative", func(t *testing.T) {
		if got := Frames(-1); len(got) != 0 {
			t.Errorf("Frames(-1) = %v, want empty", got)
		}
	})

	t.Run("names_caller_first", func(t *testing.T) {
		got := Frames(3)
		if len(got) == 0 {
			t.Fatal("Frames(3) returned nothing")
		}
		if len(got) > 3 {
			t.Fatalf("len = %d, want at most 3", len(got))
		}
		if !strings.Contains(got[0], "TestFrames") {
			t.Errorf("got[0] = %q, want it to name the calling test function", got[0])
		}
	})

	t.Run("respects_limit", func(t *testing.T) {
		if got := Frames(1); len(got) != 1 {
			t.Errorf("len(Frames(1)) = %d, want 1", len(got))
		}
	})
}

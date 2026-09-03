package wrapall

import (
	"errors"
	"testing"
)

func TestWrapAll(t *testing.T) {
	t.Run("nil_slice", func(t *testing.T) {
		if got := WrapAll(nil); got != nil {
			t.Errorf("WrapAll(nil) = %v, want nil", got)
		}
	})

	t.Run("all_nil", func(t *testing.T) {
		if got := WrapAll([]error{nil, nil}); got != nil {
			t.Errorf("WrapAll = %v, want nil", got)
		}
	})

	t.Run("keeps_original_index", func(t *testing.T) {
		got := WrapAll([]error{nil, ErrJob, nil, ErrJob})
		if len(got) != 2 {
			t.Fatalf("len = %d, want 2", len(got))
		}
		if got[0].Error() != "job 1: job failed" {
			t.Errorf("got[0] = %q, want %q", got[0].Error(), "job 1: job failed")
		}
		if got[1].Error() != "job 3: job failed" {
			t.Errorf("got[1] = %q, want %q", got[1].Error(), "job 3: job failed")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		got := WrapAll([]error{ErrJob})
		if !errors.Is(got[0], ErrJob) {
			t.Error("errors.Is = false, want true")
		}
	})
}

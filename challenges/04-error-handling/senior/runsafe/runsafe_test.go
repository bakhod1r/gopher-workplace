package runsafe

import (
	"errors"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	errStep := errors.New("step failed")

	t.Run("success", func(t *testing.T) {
		if err := Run(func() error { return nil }); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("returned_error_passthrough", func(t *testing.T) {
		err := Run(func() error { return errStep })
		if !errors.Is(err, errStep) {
			t.Fatalf("err = %v, want errStep", err)
		}
		if errors.Is(err, ErrRuntime) {
			t.Error("a returned error must not be reported as a panic")
		}
	})

	t.Run("index_panic", func(t *testing.T) {
		err := Run(func() error {
			s := []int{}
			_ = s[3]
			return nil
		})
		if !errors.Is(err, ErrRuntime) {
			t.Fatalf("err = %v, want it to match ErrRuntime", err)
		}
		if !strings.Contains(err.Error(), "index out of range") {
			t.Errorf("message = %q, want it to mention the runtime message", err.Error())
		}
	})

	t.Run("nil_map_panic", func(t *testing.T) {
		err := Run(func() error {
			var m map[string]int
			m["a"] = 1
			return nil
		})
		if !errors.Is(err, ErrRuntime) {
			t.Errorf("err = %v, want it to match ErrRuntime", err)
		}
	})
}

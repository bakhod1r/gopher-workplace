package namedresult

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v, err := Do("load", func() (int, error) { return 7, nil })
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if v != 7 {
			t.Errorf("v = %d, want 7", v)
		}
	})

	t.Run("annotates", func(t *testing.T) {
		_, err := Do("load", func() (int, error) { return 0, ErrBoom })
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if err.Error() != "load: boom" {
			t.Errorf("message = %q, want %q", err.Error(), "load: boom")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		_, err := Do("load", func() (int, error) { return 0, ErrBoom })
		if !errors.Is(err, ErrBoom) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("keeps_value_on_failure", func(t *testing.T) {
		v, _ := Do("load", func() (int, error) { return 3, ErrBoom })
		if v != 3 {
			t.Errorf("v = %d, want the value f returned (3)", v)
		}
	})
}

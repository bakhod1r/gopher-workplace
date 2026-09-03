package mustvalue

import "testing"

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if got := Must(42, nil); got != 42 {
			t.Errorf("Must(42, nil) = %d, want 42", got)
		}
	})

	t.Run("zero_value_success", func(t *testing.T) {
		if got := Must(0, nil); got != 0 {
			t.Errorf("Must(0, nil) = %d, want 0", got)
		}
	})

	t.Run("panics_with_error", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("Must did not panic")
			}
			if r != error(ErrLoad) {
				t.Errorf("panic value = %v (%T), want ErrLoad", r, r)
			}
		}()
		Must(0, ErrLoad)
	})
}

package deferorder

import (
	"errors"
	"reflect"
	"testing"
)

func TestCloseAll(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if err := CloseAll(); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("reverse_order", func(t *testing.T) {
		var order []string
		err := CloseAll(
			func() error { order = append(order, "a"); return nil },
			func() error { order = append(order, "b"); return nil },
			func() error { order = append(order, "c"); return nil },
		)
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		want := []string{"c", "b", "a"}
		if !reflect.DeepEqual(order, want) {
			t.Errorf("order = %v, want %v", order, want)
		}
	})

	t.Run("continues_past_failure", func(t *testing.T) {
		errA := errors.New("a failed")
		errC := errors.New("c failed")
		ran := 0
		err := CloseAll(
			func() error { ran++; return errA },
			func() error { ran++; return nil },
			func() error { ran++; return errC },
		)
		if ran != 3 {
			t.Fatalf("ran = %d, want 3", ran)
		}
		if !errors.Is(err, errA) || !errors.Is(err, errC) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}

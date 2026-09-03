package panicstack

import (
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		if err := Trace(func() {}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("captures_value", func(t *testing.T) {
		err := Trace(func() { panic("boom") })
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if !strings.Contains(err.Error(), "boom") {
			t.Errorf("message = %q, want it to contain the panic value", err.Error())
		}
	})

	t.Run("captures_stack", func(t *testing.T) {
		err := Trace(func() { panic("boom") })
		if !strings.Contains(err.Error(), "panicstack.") {
			t.Errorf("message = %q, want it to contain a stack frame from this package", err.Error())
		}
	})

	t.Run("int_payload", func(t *testing.T) {
		err := Trace(func() { panic(7) })
		if err == nil || !strings.Contains(err.Error(), "7") {
			t.Errorf("err = %v, want it to contain the payload", err)
		}
	})
}

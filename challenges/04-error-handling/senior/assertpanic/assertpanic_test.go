package assertpanic

import "testing"

func TestPanicked(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		v, ok := Panicked(func() {})
		if ok || v != nil {
			t.Errorf("Panicked = %v, %v, want nil, false", v, ok)
		}
	})

	t.Run("string_payload", func(t *testing.T) {
		v, ok := Panicked(func() { panic("x") })
		if !ok || v != "x" {
			t.Errorf("Panicked = %v, %v, want \"x\", true", v, ok)
		}
	})

	t.Run("int_payload", func(t *testing.T) {
		v, ok := Panicked(func() { panic(7) })
		if !ok || v != 7 {
			t.Errorf("Panicked = %v, %v, want 7, true", v, ok)
		}
	})

	t.Run("nil_payload", func(t *testing.T) {
		v, ok := Panicked(func() { panic(nil) })
		if !ok {
			t.Fatal("Panicked reported no panic for panic(nil)")
		}
		if v == nil {
			t.Error("value = nil, want the runtime's substituted panic value")
		}
	})

	t.Run("does_not_escape", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic escaped: %v", r)
			}
		}()
		Panicked(func() { panic("x") })
	})
}

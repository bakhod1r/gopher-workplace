package repanic

import (
	"runtime"
	"testing"
)

func TestHandle(t *testing.T) {
	t.Run("no_panic", func(t *testing.T) {
		if err := Handle(func() {}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("application_panic_becomes_error", func(t *testing.T) {
		err := Handle(func() { panic("app bug") })
		if err == nil {
			t.Fatal("err = nil, want an error")
		}
		if err.Error() != "panic: app bug" {
			t.Errorf("message = %q, want %q", err.Error(), "panic: app bug")
		}
	})

	t.Run("runtime_panic_propagates", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("runtime panic was absorbed, want it to propagate")
			}
			if _, ok := r.(runtime.Error); !ok {
				t.Errorf("payload %v (%T), want a runtime.Error", r, r)
			}
		}()
		Handle(func() {
			s := []int{}
			_ = s[3]
		})
	})

	t.Run("error_payload", func(t *testing.T) {
		err := Handle(func() { panic(7) })
		if err == nil || err.Error() != "panic: 7" {
			t.Errorf("err = %v, want %q", err, "panic: 7")
		}
	})
}

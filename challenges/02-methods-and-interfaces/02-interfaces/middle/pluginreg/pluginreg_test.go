package pluginreg

import (
	"errors"
	"testing"
)

func TestRegisterDuplicate(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(&Simple{N: "a"}); err != nil {
		t.Fatalf("first Register = %v", err)
	}
	if err := r.Register(&Simple{N: "a"}); !errors.Is(err, ErrDuplicate) {
		t.Errorf("second Register = %v, want ErrDuplicate", err)
	}
	if n := len(r.RunAll()); n != 1 {
		t.Errorf("RunAll len = %d, want 1", n)
	}
}

func TestRunAllOrder(t *testing.T) {
	r := NewRegistry()
	r.Register(&Simple{N: "b"})
	r.Register(&Closeable{N: "a"})

	got := r.RunAll()
	want := []string{"run:b", "run:a"}
	if len(got) != len(want) {
		t.Fatalf("RunAll = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("RunAll = %v, want %v", got, want)
		}
	}
}

func TestCloseAll(t *testing.T) {
	r := NewRegistry()
	s := &Simple{N: "s"}
	c := &Closeable{N: "c"}
	r.Register(s)
	r.Register(c)

	if got := r.CloseAll(); got != 1 {
		t.Errorf("CloseAll = %d, want 1", got)
	}
	if !c.Closed {
		t.Error("Closeable was not closed")
	}

	if got := NewRegistry().CloseAll(); got != 0 {
		t.Errorf("empty CloseAll = %d, want 0", got)
	}
}

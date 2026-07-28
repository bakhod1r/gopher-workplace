package recovertype

import (
	"errors"
	"testing"
)

func TestCall(t *testing.T) {
	sentinel := errors.New("boom")
	if got := Call(func() { panic(sentinel) }); got != sentinel {
		t.Errorf("=%v want the panicked error", got)
	}
	if got := Call(func() { panic("just a string") }); got != nil {
		t.Errorf("non-error panic should give nil, got %v", got)
	}
	if got := Call(func() {}); got != nil {
		t.Errorf("normal call should give nil")
	}
}

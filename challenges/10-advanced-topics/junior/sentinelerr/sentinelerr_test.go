package sentinelerr

import (
	"errors"
	"testing"
)

var sink error

func TestValidate(t *testing.T) {
	if err := Validate(5); err != nil {
		t.Errorf("Validate(5) = %v, want nil", err)
	}
	if err := Validate(-1); !errors.Is(err, ErrNegative) {
		t.Errorf("Validate(-1) = %v, want ErrNegative", err)
	}
	if err := Validate(MaxCount + 1); !errors.Is(err, ErrTooLarge) {
		t.Errorf("Validate = %v, want ErrTooLarge", err)
	}
	if err := Validate(0); err != nil {
		t.Errorf("Validate(0) = %v, want nil", err)
	}
	if err := Validate(MaxCount); err != nil {
		t.Errorf("Validate(MaxCount) = %v, want nil", err)
	}
}

func TestValidateAllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(200, func() { sink = Validate(-1) }); n != 0 {
		t.Errorf("Validate made %v allocations, want 0: return the sentinel, do not format one", n)
	}
}

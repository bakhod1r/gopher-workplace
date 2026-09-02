package erriface

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	e := &ValidationError{Field: "name", Message: "required"}
	if got := e.Error(); got != "name: required" {
		t.Errorf("Error() = %q, want \"name: required\"", got)
	}
}

func TestValidate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		err := Validate("")
		if err == nil {
			t.Fatal("Validate(\"\") = nil, want error")
		}
		var ve *ValidationError
		if !errors.As(err, &ve) {
			t.Fatalf("err is %T, want *ValidationError", err)
		}
		if ve.Field != "name" || ve.Message != "required" {
			t.Errorf("got %+v", *ve)
		}
	})

	t.Run("ok", func(t *testing.T) {
		if err := Validate("ann"); err != nil {
			t.Errorf("Validate(\"ann\") = %v, want nil", err)
		}
	})

	t.Run("whitespace_is_valid", func(t *testing.T) {
		if err := Validate(" "); err != nil {
			t.Errorf("Validate(\" \") = %v, want nil", err)
		}
	})
}

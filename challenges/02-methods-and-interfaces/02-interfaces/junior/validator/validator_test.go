package validator

import (
	"errors"
	"testing"
)

func TestRules(t *testing.T) {
	if err := (NotEmpty{}).Validate("a"); err != nil {
		t.Errorf("NotEmpty(\"a\") = %v, want nil", err)
	}
	if err := (NotEmpty{}).Validate(""); !errors.Is(err, ErrEmpty) {
		t.Errorf("NotEmpty(\"\") = %v, want ErrEmpty", err)
	}
	if err := (MaxLen{N: 2}).Validate("ab"); err != nil {
		t.Errorf("MaxLen exact = %v, want nil", err)
	}
	if err := (MaxLen{N: 2}).Validate("abc"); !errors.Is(err, ErrTooLong) {
		t.Errorf("MaxLen over = %v, want ErrTooLong", err)
	}
}

func TestValidateAll(t *testing.T) {
	rules := []Validator{NotEmpty{}, MaxLen{N: 2}}

	if err := ValidateAll(rules, "ab"); err != nil {
		t.Errorf("ValidateAll = %v, want nil", err)
	}
	if err := ValidateAll(rules, ""); !errors.Is(err, ErrEmpty) {
		t.Errorf("ValidateAll = %v, want ErrEmpty", err)
	}
	if err := ValidateAll(rules, "abc"); !errors.Is(err, ErrTooLong) {
		t.Errorf("ValidateAll = %v, want ErrTooLong", err)
	}
	if err := ValidateAll(nil, "anything"); err != nil {
		t.Errorf("ValidateAll(nil) = %v, want nil", err)
	}
}

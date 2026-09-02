package niliface

import "testing"

func TestRunSuccessIsNil(t *testing.T) {
	err := Run(false)
	if err != nil {
		t.Fatalf("Run(false) = %v (%T), want nil interface", err, err)
	}
	if !IsNil(err) {
		t.Error("IsNil(Run(false)) = false, want true")
	}
}

func TestRunFailure(t *testing.T) {
	err := Run(true)
	if err == nil {
		t.Fatal("Run(true) = nil, want error")
	}
	if got := err.Error(); got != "op failed" {
		t.Errorf("Error() = %q, want \"op failed\"", got)
	}
	if IsNil(err) {
		t.Error("IsNil on a real error = true")
	}
}

func TestTypedNilIsNotNil(t *testing.T) {
	var typed *OpError
	var iface error = typed
	if IsNil(iface) {
		t.Error("a typed nil pointer in an interface must not compare equal to nil")
	}
}

func TestFailedCount(t *testing.T) {
	if got := FailedCount([]error{Run(false), Run(true), Run(false)}); got != 1 {
		t.Errorf("FailedCount = %d, want 1", got)
	}
	if got := FailedCount(nil); got != 0 {
		t.Errorf("FailedCount(nil) = %d, want 0", got)
	}
}

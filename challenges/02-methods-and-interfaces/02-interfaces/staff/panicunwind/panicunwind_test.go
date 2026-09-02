package panicunwind

import (
	"errors"
	"strings"
	"testing"
)

var errNormal = errors.New("normal failure")

func TestSafeRunCatchesPanic(t *testing.T) {
	err := SafeRun(TaskFunc(func() error {
		panic("boom")
	}))
	if err == nil {
		t.Fatal("err = nil, want an error from the panic")
	}
	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("err = %v, want it to mention the panic value", err)
	}
}

func TestSafeRunPassesErrorsThrough(t *testing.T) {
	err := SafeRun(TaskFunc(func() error { return errNormal }))
	if !errors.Is(err, errNormal) {
		t.Errorf("err = %v, want the task's own error", err)
	}
}

func TestSafeRunSuccess(t *testing.T) {
	if err := SafeRun(TaskFunc(func() error { return nil })); err != nil {
		t.Errorf("err = %v, want nil", err)
	}
}

func TestSafeRunRuntimeError(t *testing.T) {
	err := SafeRun(TaskFunc(func() error {
		var s []int
		_ = s[5] // index out of range
		return nil
	}))
	if err == nil {
		t.Fatal("err = nil, want a recovered runtime error")
	}
	if !strings.Contains(err.Error(), "index out of range") {
		t.Errorf("err = %v, want the runtime error message preserved", err)
	}
}

func TestSafeRunNilDereference(t *testing.T) {
	err := SafeRun(TaskFunc(func() error {
		var p *int
		_ = *p
		return nil
	}))
	if err == nil {
		t.Fatal("err = nil, want a recovered nil dereference")
	}
}

func TestRunAll(t *testing.T) {
	errs := RunAll([]Task{
		TaskFunc(func() error { return nil }),
		TaskFunc(func() error { panic("boom") }),
		TaskFunc(func() error { return errNormal }),
	})

	if len(errs) != 3 {
		t.Fatalf("len = %d, want 3", len(errs))
	}
	if errs[0] != nil {
		t.Errorf("errs[0] = %v, want nil", errs[0])
	}
	if errs[1] == nil || !strings.Contains(errs[1].Error(), "boom") {
		t.Errorf("errs[1] = %v, want the panic error", errs[1])
	}
	if !errors.Is(errs[2], errNormal) {
		t.Errorf("errs[2] = %v, want errNormal", errs[2])
	}

	if n := len(RunAll(nil)); n != 0 {
		t.Errorf("RunAll(nil) len = %d, want 0", n)
	}
}

func TestRunAllContinuesAfterPanic(t *testing.T) {
	ran := 0
	RunAll([]Task{
		TaskFunc(func() error { ran++; panic("a") }),
		TaskFunc(func() error { ran++; panic("b") }),
		TaskFunc(func() error { ran++; return nil }),
	})
	if ran != 3 {
		t.Errorf("%d tasks ran, want 3", ran)
	}
}

func TestDeferOrderIsLIFO(t *testing.T) {
	got := Order()
	want := []string{"first", "second", "third"}

	if len(got) != len(want) {
		t.Fatalf("Order = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Order = %v, want %v", got, want)
		}
	}
}

func TestOrderDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Order let a panic escape: %v", r)
		}
	}()
	Order()
}

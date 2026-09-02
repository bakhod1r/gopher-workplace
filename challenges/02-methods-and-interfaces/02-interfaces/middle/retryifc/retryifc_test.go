package retryifc

import (
	"errors"
	"testing"
)

func TestRetrySucceedsAfterFailures(t *testing.T) {
	f := &Flaky{FailTimes: 2, Value: "ok"}
	got, err := Retry(f, 3)
	if err != nil || got != "ok" {
		t.Fatalf("Retry = %q, %v; want \"ok\", nil", got, err)
	}
	if f.Calls != 3 {
		t.Errorf("Calls = %d, want 3", f.Calls)
	}
}

func TestRetryExhausts(t *testing.T) {
	f := &Flaky{FailTimes: 5}
	got, err := Retry(f, 2)
	if !errors.Is(err, ErrTemporary) {
		t.Fatalf("err = %v, want ErrTemporary", err)
	}
	if got != "" {
		t.Errorf("got = %q, want empty", got)
	}
	if f.Calls != 2 {
		t.Errorf("Calls = %d, want 2", f.Calls)
	}
}

func TestRetryStopsOnFatal(t *testing.T) {
	p := &Permanent{}
	_, err := Retry(p, 5)
	if !errors.Is(err, ErrFatal) {
		t.Fatalf("err = %v, want ErrFatal", err)
	}
	if p.Calls != 1 {
		t.Errorf("Calls = %d, want 1 (fatal must not retry)", p.Calls)
	}
}

func TestRetryFirstTry(t *testing.T) {
	f := &Flaky{FailTimes: 0, Value: "v"}
	got, err := Retry(f, 3)
	if err != nil || got != "v" {
		t.Errorf("Retry = %q, %v", got, err)
	}
	if f.Calls != 1 {
		t.Errorf("Calls = %d, want 1", f.Calls)
	}
}

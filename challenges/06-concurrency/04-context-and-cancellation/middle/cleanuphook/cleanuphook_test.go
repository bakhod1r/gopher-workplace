package cleanuphook

import (
	"context"
	"testing"
)

func TestContextTeardown(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	h := Register(ctx)
	cancel()

	if got := h.Wait(); got != "context" {
		t.Errorf("Wait() = %q, want %q", got, "context")
	}
}

func TestReleaseBeatsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := Register(ctx)
	if !h.Release() {
		t.Error("Release() = false, want true — the hook had not run yet")
	}
	if got := h.Wait(); got != "release" {
		t.Errorf("Wait() = %q, want %q", got, "release")
	}
}

func TestReleaseAfterContextReportsFalse(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	h := Register(ctx)
	cancel()
	h.Wait() // the hook has definitely run by now

	if h.Release() {
		t.Error("Release() = true, want false — the context already tore it down")
	}
	if got := h.Wait(); got != "context" {
		t.Errorf("Wait() = %q, want %q", got, "context")
	}
}

func TestAlreadyFinishedContextFiresImmediately(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	h := Register(ctx)
	if got := h.Wait(); got != "context" {
		t.Errorf("Wait() = %q, want %q", got, "context")
	}
}

func TestDeadlineTeardown(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	h := Register(ctx)
	if got := h.Wait(); got != "context" {
		t.Errorf("Wait() = %q, want %q", got, "context")
	}
}

func TestSecondReleaseIsFalse(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := Register(ctx)
	if !h.Release() {
		t.Fatal("first Release() = false, want true")
	}
	if h.Release() {
		t.Error("second Release() = true, want false")
	}
}

package stepbudget

import (
	"context"
	"errors"
	"testing"
)

var errBoom = errors.New("upstream refused")

func ok(context.Context) error { return nil }

func bad(context.Context) error { return errBoom }

func TestRunSteps(t *testing.T) {
	cases := []struct {
		name      string
		steps     []Step
		wantRan   int
		wantError error
	}{
		{"all_succeed", []Step{ok, ok}, 2, nil},
		{"stops_at_error", []Step{ok, bad, ok}, 2, errBoom},
		{"first_fails", []Step{bad, ok}, 1, errBoom},
		{"empty_chain", nil, 0, nil},
		{"single_step", []Step{ok}, 1, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ran, err := RunSteps(context.Background(), tc.steps)
			if ran != tc.wantRan || !errors.Is(err, tc.wantError) {
				t.Errorf("RunSteps() = %d, %v; want %d, %v", ran, err, tc.wantRan, tc.wantError)
			}
		})
	}
}

func TestCancelledParentRunsNothing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	ran, err := RunSteps(ctx, []Step{ok, ok})
	if ran != 0 || !errors.Is(err, context.Canceled) {
		t.Errorf("RunSteps() = %d, %v; want 0, context.Canceled", ran, err)
	}
}

func TestExpiredDeadlineRunsNothing(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	ran, err := RunSteps(ctx, []Step{ok})
	if ran != 0 || !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("RunSteps() = %d, %v; want 0, context.DeadlineExceeded", ran, err)
	}
}

func TestEachStepGetsALiveChildContext(t *testing.T) {
	var seen []context.Context

	record := func(ctx context.Context) error {
		if err := ctx.Err(); err != nil {
			t.Errorf("step received an already-finished context: %v", err)
		}
		seen = append(seen, ctx)
		return nil
	}

	ran, err := RunSteps(context.Background(), []Step{record, record})
	if ran != 2 || err != nil {
		t.Fatalf("RunSteps() = %d, %v; want 2, nil", ran, err)
	}
	if seen[0] == seen[1] {
		t.Error("both steps shared one context; each step needs its own")
	}
	for i, ctx := range seen {
		if ctx.Err() == nil {
			t.Errorf("step %d's context was not cancelled after it returned", i)
		}
	}
}

func TestParentCancelledMidChain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	steps := []Step{
		func(context.Context) error { cancel(); return nil },
		ok,
	}

	ran, err := RunSteps(ctx, steps)
	if ran != 1 || !errors.Is(err, context.Canceled) {
		t.Errorf("RunSteps() = %d, %v; want 1, context.Canceled", ran, err)
	}
}

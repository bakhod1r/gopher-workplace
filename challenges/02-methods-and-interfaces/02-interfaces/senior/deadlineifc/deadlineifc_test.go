package deadlineifc

import (
	"testing"
	"time"
)

func TestPassedDeadlineRunsNothing(t *testing.T) {
	c := &FakeClock{T: time.Unix(100, 0)}
	op := &CountingOp{Clock: c, Cost: time.Second}

	if got := RunUntil(c, time.Unix(100, 0), op); got != 0 {
		t.Errorf("RunUntil = %d, want 0", got)
	}
	if op.Runs != 0 {
		t.Errorf("Runs = %d, want 0", op.Runs)
	}
}

func TestRunsWithinBudget(t *testing.T) {
	c := &FakeClock{T: time.Unix(0, 0)}
	op := &CountingOp{Clock: c, Cost: time.Second}

	got := RunUntil(c, time.Unix(5, 0), op)
	if got != 5 {
		t.Errorf("RunUntil = %d, want 5", got)
	}
	if op.Runs != 5 {
		t.Errorf("Runs = %d, want 5", op.Runs)
	}
}

func TestStopsExactlyAtDeadline(t *testing.T) {
	c := &FakeClock{T: time.Unix(0, 0)}
	op := &CountingOp{Clock: c, Cost: 2 * time.Second}

	got := RunUntil(c, time.Unix(5, 0), op)
	if got != 3 {
		t.Errorf("RunUntil = %d, want 3 (at t=6 the deadline has passed)", got)
	}
	if !c.T.Equal(time.Unix(6, 0)) {
		t.Errorf("clock = %v, want t=6", c.T.Unix())
	}
}

func TestSingleOperationFits(t *testing.T) {
	c := &FakeClock{T: time.Unix(0, 0)}
	op := &CountingOp{Clock: c, Cost: time.Hour}

	if got := RunUntil(c, time.Unix(1, 0), op); got != 1 {
		t.Errorf("RunUntil = %d, want 1", got)
	}
}

func TestDeadlineInThePast(t *testing.T) {
	c := &FakeClock{T: time.Unix(100, 0)}
	op := &CountingOp{Clock: c, Cost: time.Second}

	if got := RunUntil(c, time.Unix(50, 0), op); got != 0 {
		t.Errorf("RunUntil = %d, want 0", got)
	}
}

func TestIsOp(t *testing.T) {
	c := &FakeClock{T: time.Unix(0, 0)}
	var op Op = &CountingOp{Clock: c, Cost: time.Second}
	op.Do()
	if !c.T.Equal(time.Unix(1, 0)) {
		t.Errorf("clock = %d, want 1", c.T.Unix())
	}
}

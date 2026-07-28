// Package frame computes a frame time budget from a target FPS.
package frame

// TargetFPS is the frames-per-second goal.
const TargetFPS = 60

// FrameBudgetMicros returns the per-frame budget in microseconds
// (1_000_000 / TargetFPS), using integer constant division.
//
// TODO(candidate): implement — mind integer division truncation.
func FrameBudgetMicros() int {
	panic("not implemented")
}

// OverBudget reports whether a frame taking us microseconds missed the budget.
//
// TODO(candidate): implement.
func OverBudget(us int) bool {
	panic("not implemented")
}

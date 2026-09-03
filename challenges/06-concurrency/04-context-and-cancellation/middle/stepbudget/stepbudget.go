// Package stepbudget — Gopher Workplace challenge.
package stepbudget

import "context"

// Step is one upstream call in the checkout chain.
type Step func(ctx context.Context) error

// RunSteps runs the steps in order, giving each one its own cancellable
// sub-context that is cancelled before the next step starts. It stops at the
// first error and returns it, and it checks the parent context before every
// step so a cancelled request does no further work.
//
// It returns the number of steps that ran and the error that stopped the
// chain (nil when all steps succeeded).
//
// Examples:
//
//	RunSteps(live ctx, [ok, ok])        => 2, nil
//	RunSteps(live ctx, [ok, bad, ok])   => 2, the bad step's error
//	RunSteps(cancelled ctx, [ok])       => 0, context.Canceled
func RunSteps(ctx context.Context, steps []Step) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Package retryifc — Gopher Workplace challenge.
package retryifc

import "errors"

// Failure kinds.
var (
	ErrTemporary = errors.New("temporary")
	ErrFatal     = errors.New("fatal")
)

// Op is a retryable operation.
type Op interface {
	Do() (string, error)
}

// Flaky fails FailTimes times, then succeeds.
type Flaky struct {
	FailTimes int
	Value     string
	Calls     int
}

// Do fails until FailTimes calls have been made.
func (f *Flaky) Do() (string, error) {
	// TODO(candidate): count the call; fail while Calls <= FailTimes.
	panic("not implemented")
}

// Permanent always fails fatally.
type Permanent struct {
	Calls int
}

// Do always returns ErrFatal.
func (p *Permanent) Do() (string, error) {
	// TODO(candidate): count the call, return ErrFatal.
	panic("not implemented")
}

// Retry calls op up to attempts times, stopping at the first success or the
// first non-temporary error. It returns the last error when all attempts fail.
//
// Examples:
//
//	Retry(&Flaky{FailTimes: 2, Value: "ok"}, 3) => "ok", nil
//	Retry(&Flaky{FailTimes: 5}, 2)              => "", ErrTemporary
func Retry(op Op, attempts int) (string, error) {
	// TODO(candidate): bounded retry loop.
	panic("not implemented")
}

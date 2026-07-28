// Package epsiloneq compares floats within a tolerance.
package epsiloneq

// Equal reports whether a and b are within eps of each other (absolute
// tolerance). NaN is never equal to anything.
//
// TODO(candidate): implement using math.Abs, guarding against NaN.
func Equal(a, b, eps float64) bool {
	panic("not implemented")
}

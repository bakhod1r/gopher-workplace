// Package tempconv — Gopher Workplace challenge.
package tempconv

// ToFloat converts any float-like named type to float64.
func ToFloat[T ~float64](v T) float64 {
	// TODO(candidate): convert to float64.
	panic("not implemented")
}

// FromFloat converts a float64 into a float-like named type.
func FromFloat[T ~float64](f float64) T {
	// TODO(candidate): convert from float64.
	panic("not implemented")
}

// Rescale converts v through f, keeping the named type.
func Rescale[T ~float64](v T, f func(float64) float64) T {
	// TODO(candidate): convert out, apply f, convert back.
	panic("not implemented")
}

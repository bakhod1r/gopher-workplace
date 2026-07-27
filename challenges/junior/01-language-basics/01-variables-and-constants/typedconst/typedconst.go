// Package typedconst — Gopher Workplace challenge.
package typedconst

// TODO(candidate): declare the two batching constants.
//
//	MaxBatch — a *typed* constant of type byte, value 200. Writing the type
//	           pins it: it is a byte everywhere it appears, so mixing it with
//	           an int needs an explicit conversion.
//	Retries  — an *untyped* constant, value 3. With no type written down it
//	           adapts to whatever context it lands in — int here, float64
//	           there — with no conversion at all.
//
// Declare MaxBatch with an explicit byte type and Retries with none.
const (
	MaxBatch byte = 0
	Retries       = 0
)

// Fits reports whether n items fit in a single batch, i.e. whether n is at most
// MaxBatch. Negative n never fits.
//
// MaxBatch is a byte and n is an int, so this comparison needs a conversion.
//
// Examples:
//
//	Fits(0)   => true
//	Fits(200) => true
//	Fits(201) => false
//	Fits(-1)  => false
func Fits(n int) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Budget returns the total time budget for one item: base seconds per attempt,
// across Retries attempts.
//
// Retries is untyped, so it multiplies a float64 directly — no conversion.
//
// Examples:
//
//	Budget(1.5) => 4.5
//	Budget(0)   => 0
//	Budget(0.1) => 0.30000000000000004
func Budget(base float64) float64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

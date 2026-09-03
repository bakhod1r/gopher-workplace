// Package histmerge — Gopher Workplace challenge.
package histmerge

// Hist is a histogram: ascending upper bounds and one count per bound plus a
// final overflow count, so len(Counts) == len(Bounds)+1.
type Hist struct {
	Bounds []float64
	Counts []int64
}

// Valid reports whether h is well formed: bounds strictly ascending and the
// counts one longer than the bounds.
//
// Examples:
//
//	Valid(Hist{[]float64{1, 2}, []int64{0, 0, 0}}) => true
func Valid(h Hist) bool {
	panic("not implemented")
}

// Merge adds two histograms that share the same bounds, summing their counts.
// Histograms with different bounds cannot be merged — counts under different
// bucket edges mean different things — so that reports false. Neither input is
// modified.
//
// Examples:
//
//	Merge(a, b) => summed histogram, true
func Merge(a, b Hist) (Hist, bool) {
	panic("not implemented")
}

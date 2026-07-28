// Package area multiplies two 32-bit dimensions. Planted narrow intermediate wraps.
package area

// Area returns w*h without overflow for 32-bit inputs.
func Area(w, h int32) int64 {
	// CHANGE CODE BELOW THIS LINE
	return int64(w * h)
	// CHANGE CODE ABOVE THIS LINE
}

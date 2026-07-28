// Package stock tracks inventory. A planted unsigned-underflow bug lurks.
package stock

// Remaining returns how many units are left after selling sold from have.
// It must never report more than were on hand; overselling yields 0.
func Remaining(have, sold uint) uint {
	// CHANGE CODE BELOW THIS LINE
	return have - sold
	// CHANGE CODE ABOVE THIS LINE
}

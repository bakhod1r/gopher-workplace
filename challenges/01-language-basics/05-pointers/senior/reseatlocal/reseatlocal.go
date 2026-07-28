// Package reseatlocal reseats a caller's pointer via a double pointer. A planted
// bug assigns to the local parameter instead of dereferencing, so the caller's
// pointer never changes.
package reseatlocal

// Reseat should make the caller's pointer (*pp) point to q.
func Reseat(pp **int, q *int) {
	// CHANGE CODE BELOW THIS LINE
	pp = &q
	// CHANGE CODE ABOVE THIS LINE
}

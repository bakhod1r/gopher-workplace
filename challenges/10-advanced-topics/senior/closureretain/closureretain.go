// Package closureretain — Gopher Workplace challenge.
package closureretain

// Record is one ingested item.
type Record struct {
	Size int
	Pad  [256]byte
}

// Summarize returns a function reporting the batch's total size.
//
// The returned function outlives the batch, so it must capture the answer
// rather than the data: a closure over the slice keeps every record alive
// for as long as the callback exists.
//
// Examples:
//
//	f := Summarize(batch); f() => the total size
func Summarize(batch []Record) func() int {
	// CHANGE CODE BELOW THIS LINE
	return func() int {
		total := 0
		for _, r := range batch {
			total += r.Size
		}
		return total
	}
	// CHANGE CODE ABOVE THIS LINE
}

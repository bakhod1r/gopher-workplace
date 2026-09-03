// Package filterretain — Gopher Workplace challenge.
package filterretain

// Record is one ingested item.
type Record struct {
	ID   int
	Size int
	Pad  [64]byte
}

// Keep returns the records whose Size is at least min.
//
// Typical batches are huge and typical results are tiny. The result must
// not keep the batch's storage alive once the caller drops the batch.
//
// Examples:
//
//	Keep(batch, 100) => only the large records
func Keep(records []Record, min int) []Record {
	// CHANGE CODE BELOW THIS LINE
	k := 0
	for _, r := range records {
		if r.Size >= min {
			records[k] = r
			k++
		}
	}
	return records[:k]
	// CHANGE CODE ABOVE THIS LINE
}

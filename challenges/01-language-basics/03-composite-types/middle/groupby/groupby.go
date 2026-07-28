// Package groupby groups strings by their first letter.
package groupby

// ByFirst groups words by their first byte (assumed ASCII). Order within each
// group follows input order. Empty words are skipped.
//
// TODO(candidate): append each word to its group slice.
func ByFirst(words []string) map[byte][]string {
	panic("not implemented")
}

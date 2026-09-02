// Package zerocopy — Gopher Workplace challenge.
package zerocopy

// Parser splits a buffer into fields.
type Parser interface {
	Fields(data []byte) [][]byte
}

// ZeroCopyParser splits on a separator without copying field bytes.
type ZeroCopyParser struct {
	Sep byte
}

// Fields splits data on Sep, returning sub-slices that alias data.
//
// The returned fields share memory with data: mutating data afterwards
// changes them.
//
// Examples:
//
//	Fields([]byte("a,bb,c")) => ["a", "bb", "c"], all aliasing the input
func (p *ZeroCopyParser) Fields(data []byte) [][]byte {
	// TODO(candidate): sub-slice, do not copy field bytes.
	panic("not implemented")
}

// CopyFields splits data and returns independent copies.
func CopyFields(p Parser, data []byte) [][]byte {
	// TODO(candidate): parse, then copy each field.
	panic("not implemented")
}

// Package whennotgenericgen — Gopher Workplace challenge.
package whennotgenericgen

// Writer accepts lines of text.
type Writer interface {
	Write(line string)
}

// WriteAll writes every line to w.
// It takes an interface because it needs behaviour, not a
// type parameter, and there is only one value involved.
func WriteAll(w Writer, lines []string) int {
	// TODO(candidate): write each line, counting the calls.
	panic("not implemented")
}

// WriteEach writes values of any type by formatting them first.
// Here a type parameter earns its place: the caller keeps a
// typed slice.
func WriteEach[T any](w Writer, vs []T, format func(T) string) int {
	// TODO(candidate): format and write each value, counting the calls.
	panic("not implemented")
}

// Package byteappend — Gopher Workplace challenge.
package byteappend

// AppendRecord appends "key=value;" to dst and returns the extended slice —
// the append-to-caller's-buffer shape the standard library uses for
// strconv.AppendInt and time.Time.AppendFormat. When dst has spare capacity
// the call must not allocate at all.
//
// Examples:
//
//	AppendRecord(nil, "a", "1") => []byte("a=1;")
func AppendRecord(dst []byte, key, value string) []byte {
	panic("not implemented")
}

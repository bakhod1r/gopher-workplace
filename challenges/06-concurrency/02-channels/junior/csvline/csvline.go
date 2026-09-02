// Package csvline — Gopher Workplace challenge.
package csvline

// JoinFields drains the field stream and assembles one CSV line, placing
// sep between the fields. An empty record yields the empty string.
//
// Examples:
//
//	JoinFields(chan "a","b", ",") => "a,b"
//	JoinFields(chan "x", ",")     => "x"
//	JoinFields(closed empty, ",") => ""
func JoinFields(fields <-chan string, sep string) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

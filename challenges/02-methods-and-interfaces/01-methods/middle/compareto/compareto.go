// Package compareto — Gopher Workplace challenge.
package compareto

// Version represents a semantic version (Major.Minor only).
type Version struct {
	Major int
	Minor int
}

// Compare returns:
//
//	-1 if v < other
//	 0 if v == other
//	+1 if v > other
//
// Comparison is Major-first, then Minor.
//
// Examples:
//
//	Version{1, 0}.Compare(Version{2, 0}) => -1
//	Version{1, 2}.Compare(Version{1, 2}) => 0
//	Version{2, 0}.Compare(Version{1, 9}) => 1
func (v Version) Compare(other Version) int {
	// TODO(candidate): compare Major first, then Minor.
	panic("not implemented")
}

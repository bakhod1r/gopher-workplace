// Package nilvsempty filters a list and must return a non-nil empty slice when
// nothing matches. A planted bug returns nil.
package nilvsempty

// NonEmpty returns the non-empty strings of in. When none match, it must return
// an empty, non-nil slice (so JSON encodes [] not null).
func NonEmpty(in []string) []string {
	// CHANGE CODE BELOW THIS LINE
	var out []string
	// CHANGE CODE ABOVE THIS LINE
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

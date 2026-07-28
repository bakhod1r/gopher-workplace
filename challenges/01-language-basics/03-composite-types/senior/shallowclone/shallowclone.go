// Package shallowclone copies a record. A planted bug shares the slice field, so
// mutating the copy leaks into the original.
package shallowclone

// Doc has a title and a list of tags.
type Doc struct {
	Title string
	Tags  []string
}

// Clone returns a deep-ish copy: the Tags slice must be independent.
func Clone(d Doc) Doc {
	// CHANGE CODE BELOW THIS LINE
	return d
	// CHANGE CODE ABOVE THIS LINE
}

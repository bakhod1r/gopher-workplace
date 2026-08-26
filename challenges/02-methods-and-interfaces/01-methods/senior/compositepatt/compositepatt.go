// Package compositepatt — Gopher Workplace challenge.
package compositepatt

// Folder contains files (sizes).
type Folder struct {
	Files []int
	Sub   []*Folder
}

// Size returns total size of files and subfolders.
func (f *Folder) Size() int {
	// TODO(candidate): sum f.Files, recurse for f.Sub
	panic("not implemented")
}

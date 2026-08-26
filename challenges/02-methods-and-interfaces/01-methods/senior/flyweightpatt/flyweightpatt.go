// Package flyweightpatt — Gopher Workplace challenge.
package flyweightpatt

// FontData is heavy.
type FontData struct{ data string }

// FlyweightFactory shares FontData.
type FlyweightFactory struct {
	fonts map[string]*FontData
}

// Get returns the existing FontData or creates a new one.
func (f *FlyweightFactory) Get(name string) *FontData {
	// TODO(candidate): check if name is in fonts. If not, create & store. Return it.
	panic("not implemented")
}

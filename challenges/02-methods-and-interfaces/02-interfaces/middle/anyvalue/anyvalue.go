// Package anyvalue — Gopher Workplace challenge.
package anyvalue

// Bag holds untyped values by key.
type Bag struct {
	data map[string]any
}

// NewBag returns an empty bag.
func NewBag() *Bag {
	return &Bag{data: make(map[string]any)}
}

// Set stores a value.
func (b *Bag) Set(key string, v any) {
	// TODO(candidate): store it.
	panic("not implemented")
}

// Len returns how many keys are stored.
func (b *Bag) Len() int {
	// TODO(candidate): map length.
	panic("not implemented")
}

// GetString returns the string at key; ok is false if absent or not a string.
//
// Examples:
//
//	b.Set("a", "x"); b.GetString("a") => "x", true
//	b.Set("n", 1);   b.GetString("n") => "", false
func (b *Bag) GetString(key string) (string, bool) {
	// TODO(candidate): presence AND type must hold.
	panic("not implemented")
}

// Kinds returns the sorted distinct kind names in the bag:
// "int", "string", "bool", or "other".
func (b *Bag) Kinds() []string {
	// TODO(candidate): classify each value, dedupe, sort.
	panic("not implemented")
}

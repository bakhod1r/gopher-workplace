// Package defaultmapzerobug — Gopher Workplace challenge.
package defaultmapzerobug

// DefaultMap returns Default for keys that were never set.
type DefaultMap[K comparable, V any] struct {
	Default V
	m       map[K]V
}

// Get returns the stored value, or the default when the key is absent.
// A lookup never inserts.
func (d *DefaultMap[K, V]) Get(k K) V {
	// CHANGE CODE BELOW THIS LINE
	if d.m == nil {
		d.m = make(map[K]V)
	}
	v, ok := d.m[k]
	if !ok {
		v = d.Default
		d.m[k] = v
	}
	return v
	// CHANGE CODE ABOVE THIS LINE
}

// Set stores v under k.
func (d *DefaultMap[K, V]) Set(k K, v V) {
	if d.m == nil {
		d.m = make(map[K]V)
	}
	d.m[k] = v
}

// Len reports how many keys were explicitly set.
func (d *DefaultMap[K, V]) Len() int {
	return len(d.m)
}

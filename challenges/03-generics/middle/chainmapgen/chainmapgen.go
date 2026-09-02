// Package chainmapgen — Gopher Workplace challenge.
package chainmapgen

// Chain looks values up across layered maps, earlier layers first.
// Use NewChain to create one.
type Chain[K comparable, V any] struct {
	layers []map[K]V
}

// NewChain returns a chain over the given layers, highest
// priority first.
func NewChain[K comparable, V any](layers ...map[K]V) *Chain[K, V] {
	// TODO(candidate): store the layers in priority order.
	panic("not implemented")
}

// Get returns the value from the first layer holding k.
func (c *Chain[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): search the layers in order.
	panic("not implemented")
}

// Flatten returns one map with higher layers winning.
func (c *Chain[K, V]) Flatten() map[K]V {
	// TODO(candidate): merge the layers, letting earlier ones win.
	panic("not implemented")
}

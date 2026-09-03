// Package radixlongestbug — Gopher Workplace challenge.
package radixlongestbug

// Prefixes maps string prefixes to values, longest match wins.
type Prefixes[V any] struct {
	m map[string]V
}

// Longest returns the value registered for the longest prefix of s.
// It reports false when no registered prefix matches.
func (p *Prefixes[V]) Longest(s string) (V, bool) {
	// CHANGE CODE BELOW THIS LINE
	var best V
	for i := 0; i <= len(s); i++ {
		if v, ok := p.m[s[:i]]; ok {
			return v, true
		}
	}
	return best, false
	// CHANGE CODE ABOVE THIS LINE
}

// Add registers v under the given prefix.
func (p *Prefixes[V]) Add(prefix string, v V) {
	if p.m == nil {
		p.m = make(map[string]V)
	}
	p.m[prefix] = v
}

// Len reports how many prefixes are registered.
func (p *Prefixes[V]) Len() int {
	return len(p.m)
}

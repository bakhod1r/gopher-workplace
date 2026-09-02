// Package dsucompressbug — Gopher Workplace challenge.
package dsucompressbug

// DSU is a disjoint-set union over comparable elements.
type DSU[T comparable] struct {
	parent map[T]T
}

// Find returns the representative of x's set.
// Every node on the path is re-pointed at that representative.
func (u *DSU[T]) Find(x T) T {
	// CHANGE CODE BELOW THIS LINE
	if u.parent == nil {
		u.parent = make(map[T]T)
	}
	p, ok := u.parent[x]
	if !ok {
		u.parent[x] = x
		return x
	}
	for p != x {
		x = p
		p = u.parent[x]
	}
	return x
	// CHANGE CODE ABOVE THIS LINE
}

// Union merges the sets containing a and b.
func (u *DSU[T]) Union(a, b T) {
	ra, rb := u.Find(a), u.Find(b)
	if ra != rb {
		u.parent[ra] = rb
	}
}

// Connected reports whether a and b are in the same set.
func (u *DSU[T]) Connected(a, b T) bool {
	return u.Find(a) == u.Find(b)
}

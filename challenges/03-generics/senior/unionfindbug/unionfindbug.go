// Package unionfindbug — Gopher Workplace challenge.
package unionfindbug

// UF is a union-find over comparable elements.
type UF[T comparable] struct {
	parent map[T]T
}

// Union merges the sets containing a and b.
func (u *UF[T]) Union(a, b T) {
	// CHANGE CODE BELOW THIS LINE
	u.Find(a)
	u.Find(b)
	if a != b {
		u.parent[a] = b
	}
	// CHANGE CODE ABOVE THIS LINE
}

// Find returns the representative of x's set.
func (u *UF[T]) Find(x T) T {
	if u.parent == nil {
		u.parent = make(map[T]T)
	}
	p, ok := u.parent[x]
	if !ok {
		u.parent[x] = x
		return x
	}
	if p == x {
		return x
	}
	r := u.Find(p)
	u.parent[x] = r
	return r
}

// Connected reports whether a and b are in the same set.
func (u *UF[T]) Connected(a, b T) bool {
	return u.Find(a) == u.Find(b)
}

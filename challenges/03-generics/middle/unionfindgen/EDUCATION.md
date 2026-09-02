# Disjoint Set

## Intuition

Compression is what turns a potentially long chain into a one-hop lookup, and it is why union-find is effectively constant time in practice.

## Approach

1. `Find`: register unseen elements, stop at a self-parent, otherwise recurse and compress.
2. `Union`: point one root at the other when they differ.
3. `Connected`: compare representatives.

## Solution

```go
func NewDisjoint[T comparable]() *Disjoint[T] {
	return &Disjoint[T]{parent: make(map[T]T)}
}

func (d *Disjoint[T]) Find(v T) T {
	p, ok := d.parent[v]
	if !ok {
		d.parent[v] = v
		return v
	}
	if p == v {
		return v
	}
	root := d.Find(p)
	d.parent[v] = root
	return root
}

func (d *Disjoint[T]) Union(a, b T) {
	ra, rb := d.Find(a), d.Find(b)
	if ra != rb {
		d.parent[ra] = rb
	}
}

func (d *Disjoint[T]) Connected(a, b T) bool {
	return d.Find(a) == d.Find(b)
}
```

## Walkthrough

`Union(a,b); Union(b,c)` leaves all three sharing one root, so `Connected(a,c)` is true.

## Pitfalls

- Comparing `a` and `b` directly instead of their roots.
- Skipping compression, leaving long chains after many unions.
- Assigning `d.parent[a] = b` without resolving roots first.

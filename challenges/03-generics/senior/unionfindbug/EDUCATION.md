# Union That Links The Wrong Nodes

## Intuition

Linking `a` directly to `b` orphans `a`'s old subtree and leaves `b`'s own root untouched, so later `Find` calls land in different trees.

## Approach

1. Resolve both roots with `Find`.
2. If they differ, point one root at the other.

## Solution

```go
func (u *UF[T]) Union(a, b T) {
	ra, rb := u.Find(a), u.Find(b)
	if ra != rb {
		u.parent[ra] = rb
	}
}

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

func (u *UF[T]) Connected(a, b T) bool {
	return u.Find(a) == u.Find(b)
}
```

## Walkthrough

After `Union(1,2)` node 1 points at 2. `Union(2,3)` then sets `parent[2] = 3`, which happens to work; but `Union(3,1)` repoints 3 at 1 and rebuilds a cycle rather than merging roots.

## Pitfalls

- Calling `Find` and discarding its result.
- Comparing the arguments instead of their roots.

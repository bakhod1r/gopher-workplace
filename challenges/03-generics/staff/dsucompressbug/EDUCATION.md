# The Union-Find That Walks The Chain Forever

## Intuition

Without path compression each `Find` re-walks the entire chain, so a degenerate chain of n merges makes n queries cost O(n^2). Writing the discovered root back into every node on the path collapses the structure, and the next traversal is one hop.

## Approach

1. Return x itself when it is its own parent, or when it is new.
2. Otherwise resolve the parent's root recursively.
3. Store that root back into x before returning it.

## Solution

```go
func (u *DSU[T]) Find(x T) T {
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

func (u *DSU[T]) Union(a, b T) {
	ra, rb := u.Find(a), u.Find(b)
	if ra != rb {
		u.parent[ra] = rb
	}
}

func (u *DSU[T]) Connected(a, b T) bool {
	return u.Find(a) == u.Find(b)
}
```

## Walkthrough

After `Union(i, i+1)` for i in 0..14999 the set is a 15000-deep chain. The buggy `Find(0)` walks all 15000 links on *every* call; the fixed one walks them once and then answers in a single hop.

## Pitfalls

- Compressing only one level (`u.parent[x] = u.parent[p]`) — better, but still not flat.
- Discarding the recursive result instead of storing it.
- Assuming a passing correctness suite proves the structure is healthy.

# The Set That Panics On A Slice

## Intuition

`comparable` is a compile-time promise, and `any` is allowed to satisfy it. The moment the dynamic value inside the interface is a slice, map, or function, hashing it is a run-time panic — so the set has to ask before it indexes.

## Approach

1. Probe each element to see whether it can be hashed.
2. Pass unhashable elements straight through, undeduplicated.
3. Deduplicate the rest through the map, preserving input order.

## Solution

```go
func Distinct[T comparable](vals []T) []T {
	out := make([]T, 0, len(vals))
	seen := make(map[T]struct{}, len(vals))
	for _, v := range vals {
		if !hashable(v) {
			out = append(out, v)
			continue
		}
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func hashable[T comparable](v T) (ok bool) {
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	probe := make(map[any]struct{}, 1)
	probe[any(v)] = struct{}{}
	return true
}
```

## Walkthrough

`Distinct([]any{1, []int{2}})` reaches `seen[v]` with a `[]int` inside the interface and panics before it can return anything.

## Pitfalls

- Assuming `comparable` rules out `any` — it did before Go 1.20, and does not now.
- Recovering around the whole function, which would swallow the duplicate work as well as the panic.

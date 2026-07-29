# Deep-copying maps with reference values

## Intuition

A map copy that assigns each value copies **references** for reference-typed
values (slices, maps, pointers). To fully separate, clone the values too:

```go
for k, v := range m { out[k] = append([]int{}, v...) }
```

## Approach

1. Bug: `out[k] = v` copies the slice header, so clone and original share the same backing array; mutating the clone corrupts the source.
2. Fix: `out[k] = append([]int(nil), v...)` allocates a fresh backing array per value.

## Solution

```go
func Clone(m map[string][]int) map[string][]int {
	out := make(map[string][]int, len(m))
	for k, v := range m {
		out[k] = append([]int(nil), v...)
	}
	return out
}
```

## Walkthrough

Clone({a:[1 2]}): fix builds out[a] as a new [1 2] independent of m[a]. Then c[a][0]=99 changes only the clone; m[a][0] remains 1.

## Pitfalls

- Every reference-typed value needs its own deep copy.
- Nested maps of maps need recursion.
- Value types (int, string, array) copy fully and are safe.

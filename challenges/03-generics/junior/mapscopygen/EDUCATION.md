# Merge Maps

## Intuition

Merging is just layered copying. Allocating the destination yourself is what makes the operation non-destructive, which `maps.Copy` alone would not be.

## Approach

1. Allocate `out` sized for both inputs.
2. Copy `base` in, then `override`.
3. Return `out`.

## Solution

```go
func Merge[K comparable, V any](base, override map[K]V) map[K]V {
	out := make(map[K]V, len(base)+len(override))
	maps.Copy(out, base)
	maps.Copy(out, override)
	return out
}
```

## Walkthrough

`Merge({a:1}, {a:2})` writes `a:1` first and then overwrites it with `a:2`.

## Pitfalls

- Calling `maps.Copy(base, override)`, which mutates the caller's `base`.
- Copying `override` first, so the defaults win.
- Returning a nil map for two empty inputs.

# Reducing Slices

## Intuition

Reduction methods take a collection and return a scalar. Sum starts at 0, Product
starts at 1.

## Approach

1. `p := 1`.
2. Loop over `l` and multiply.

## Solution

```go
func (l IntList) Product() int {
	p := 1
	for _, v := range l {
		p *= v
	}
	return p
}
```

## Walkthrough

For `{2, 3, 4}`:
- `p` = 1
- `p *= 2` → 2
- `p *= 3` → 6
- `p *= 4` → 24.

## Pitfalls

- Initializing `p := 0` means the product is always 0.

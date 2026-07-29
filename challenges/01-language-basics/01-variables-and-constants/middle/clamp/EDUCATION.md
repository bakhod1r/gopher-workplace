# Block scope and shadowing

## Intuition

A variable exists only inside the block (`{ ... }`) where it is declared, and
disappears at the block's closing brace. Re-declaring a name with `:=` inside a
narrower block creates a **new** variable that *shadows* the outer one:

```go
lo, hi := 0, 10
if lo > hi {
	lo, hi := hi, lo // NEW lo, hi — the outer pair is untouched!
	_ = lo; _ = hi
}
// outer lo, hi still 0, 10
```

To modify the outer variables, use `=` (assignment), not `:=` (declaration):

```go
if lo > hi {
	lo, hi = hi, lo // updates the outer pair
}
```

## Approach

1. Normalize the bounds when `lo > hi`.
2. Clamp `v` below `lo` and above `hi`.
3. Otherwise return `v`.

## Solution

```go
func Clamp(v, lo, hi int) int {
	if lo, hi = order(lo, hi); v < lo {
		return lo
	} else if v > hi {
		return hi
	}
	return v
}

func order(lo, hi int) (int, int) {
	if lo > hi {
		return hi, lo
	}
	return lo, hi
}
```

## Walkthrough

`Clamp(5, 10, 0)` swaps to `[0, 10]`, and 5 is inside, so it returns 5.

## Pitfalls

- `:=` requires at least one new name on the left; if all names exist in the
  current scope it is a compile error — but in a *nested* block it silently
  shadows instead.
- `go vet`'s shadow analysis and linters can flag suspicious shadows.
- Multiple assignment (`lo, hi = hi, lo`) swaps without a temporary.

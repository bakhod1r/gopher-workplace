# Flyer

## Intuition

The interface normalises units. Callers get metres and never learn that a jet stores feet — the conversion lives inside the implementation.

## Approach

1. `Bird.Altitude` returns `b.Meters` unchanged.
2. `Jet.Altitude` returns `j.Feet * 3048 / 10000` — multiply first to avoid truncating to zero.
3. `Highest` keeps a `best` variable, updating it when a larger altitude appears.

## Solution

```go
func (b Bird) Altitude() int { return b.Meters }

func (j Jet) Altitude() int { return j.Feet * 3048 / 10000 }

func Highest(fs []Flyer) int {
	best := 0
	for _, f := range fs {
		if a := f.Altitude(); a > best {
			best = a
		}
	}
	return best
}
```

## Walkthrough

`Jet{Feet: 1000}`: `1000 * 3048 = 3048000`, divided by `10000` gives `304` (truncated). That beats nothing else in the slice, so `Highest` returns 304.

## Pitfalls

- `j.Feet / 10000 * 3048` — the division happens first and collapses to 0.
- Starting `best` at the first element without guarding the empty slice.
- Using floats when the signature returns `int`.

# Comparator

## Intuition

Passing the ordering as a value means one sort function serves every report. `Reverse` shows the payoff: it composes with any comparator, including another `Reverse`.

## Approach

1. `ByName` returns an explicit three-way result; `ByAge` can subtract because ages are small.
2. `Reverse` negates the inner comparison.
3. `SortWith` allocates a copy with `make` + `copy`, sorts the copy with `sort.SliceStable`, and returns it.

## Solution

```go
func (ByName) Compare(a, b Record) int {
	switch {
	case a.Name < b.Name:
		return -1
	case a.Name > b.Name:
		return 1
	default:
		return 0
	}
}

func (ByAge) Compare(a, b Record) int { return a.Age - b.Age }

func (r Reverse) Compare(a, b Record) int { return -r.Inner.Compare(a, b) }

func SortWith(recs []Record, c Comparator) []Record {
	out := make([]Record, len(recs))
	copy(out, recs)
	sort.SliceStable(out, func(i, j int) bool {
		return c.Compare(out[i], out[j]) < 0
	})
	return out
}
```

## Walkthrough

`Reverse{Inner: Reverse{Inner: ByAge{}}}` negates twice, so the order is ascending again — the test pins that down.

## Pitfalls

- Sorting `recs` directly, which mutates the caller's slice.
- `a.Age - b.Age` is fine for small ages but overflows for extreme values; the explicit three-way form is safer.
- Using `sort.Slice` where ties must stay in input order — prefer `SliceStable` when the comparator is not total.

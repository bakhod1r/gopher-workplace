# Comparable Interface

## Intuition

Sorting and ranking code does not need to know the element type — only how to order two of them. That is exactly one method.

## Approach

1. Assert `other.(Score)` to get the concrete value.
2. Return `-1`, `0`, or `1` from a `switch`.
3. `Max` returns `b` only when `b.CompareTo(a) > 0`, so ties keep `a`.

## Solution

```go
func (s Score) CompareTo(other Comparable) int {
	o := other.(Score)
	switch {
	case s < o:
		return -1
	case s > o:
		return 1
	default:
		return 0
	}
}

func Max(a, b Comparable) Comparable {
	if b.CompareTo(a) > 0 {
		return b
	}
	return a
}
```

## Walkthrough

`Max(Score(4), Score(4))`: `b.CompareTo(a)` is `0`, not `>0`, so `a` is returned — the documented tie rule.

## Pitfalls

- Returning a bool instead of a three-way int — the contract loses the equal case.
- Using `a.CompareTo(b) < 0` in `Max`, which returns `b` on ties and breaks the rule.
- Forgetting the type assertion and trying to compare a `Comparable` with `<`.

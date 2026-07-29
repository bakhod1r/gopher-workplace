# Partitioning into buckets

## Intuition

Route each element to one of two accumulators based on a predicate:

```go
for _, x := range xs {
	if x%2 == 0 { evens = append(evens, x) } else { odds = append(odds, x) }
}
```

## Approach

1. Iterate xs.
2. v%2==0 -> append to evens, else append to odds.
3. Order preserved by append.
4. Return both slices.

## Solution

```go
func Partition(xs []int) (evens, odds []int) {
	for _, v := range xs {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return evens, odds
}
```

## Walkthrough

[1..6]: 1 odd,2 even,3 odd,4 even,5 odd,6 even -> evens=[2,4,6], odds=[1,3,5].

## Pitfalls

- Named returns start as `nil`; appending is fine, and the test accepts nil/empty.
- Generalizes to N buckets with a `map[key][]T`.
- Order is preserved only within each bucket, not across them.

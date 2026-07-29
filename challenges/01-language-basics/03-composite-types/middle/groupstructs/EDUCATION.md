# Group-by-sum

## Intuition

Accumulate a running total per key; the map zero value makes `+=` work for new
keys:

```go
m := make(map[string]int)
for _, o := range orders { m[o.Customer] += o.Amount }
```

## Approach

1. Create map[string]int.
2. For each order add Amount to map[Customer].
3. The += on a missing key starts from zero.
4. Return map.

## Solution

```go
type Order struct {
	Customer string
	Amount   int
}

func TotalByCustomer(orders []Order) map[string]int {
	out := map[string]int{}
	for _, o := range orders {
		out[o.Customer] += o.Amount
	}
	return out
}
```

## Walkthrough

ann+=100 (100); bob+=50 (50); ann+=25 (125); bob+=75 (125). {ann:125,bob:125}.

## Pitfalls

- `make` the map first; `+=` is a write.
- The zero-value read is what lets a new key start at 0.
- For multiple aggregates, use a struct as the map value.

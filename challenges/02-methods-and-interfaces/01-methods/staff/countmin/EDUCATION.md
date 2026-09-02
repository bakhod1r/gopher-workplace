# Count-Min Sketch

## Intuition

A sketch is a hash table that refuses to grow: collisions are accepted, not
resolved. Because two different items sharing a bucket both increment it, any
single row's counter is an *over*-estimate. Using several independent rows and
taking the minimum keeps the error one-sided and small — the estimate is never
below the truth.

## Approach

1. `Add`: increment one counter in each row.
2. `Count`: read the same counters and return the smallest.

## Solution

```go
func (s *Sketch) Add(item string) {
	if len(item) == 0 {
		return
	}
	s.row1[h1(item)]++
	s.row2[h2(item)]++
}

func (s *Sketch) Count(item string) int {
	if len(item) == 0 {
		return 0
	}
	return min(s.row1[h1(item)], s.row2[h2(item)])
}
```

## Walkthrough

`h1` is the first byte, `h2` the last. `"apple"` maps to `('a', 'e')` and
`"bat"` to `('b', 't')` — four distinct buckets, so the two items do not
interfere. Two `Add("apple")` calls leave `row1['a'] == 2` and `row2['e'] == 2`,
so `Count("apple")` is 2. One `Add("bat")` gives 1.

Had the two items shared a bucket — as `"apple"` and `"ape"` would, both being
`('a','e')` — every estimate in that bucket would read 3, and the "want 2"
assertion would fail. That is not a bug in the code but the sketch's documented
over-estimation, which is why the test uses items with disjoint hashes.

## Pitfalls

- **Returning the max, or the sum.** Both throw away the one-sided error
  guarantee; the sum roughly doubles every count.
- **Reading only one row.** That is a plain hash counter with unbounded
  collision error.
- **Value receiver.** The rows are fixed-size arrays, so `Add` on a value
  receiver increments a copy and every count stays 0.
- **Decrementing.** Plain count-min cannot support removal; the min would drop
  below the true count and the guarantee flips.

## Why the guarantee is one-sided

Every increment for an item lands in that item's buckets, so each row's counter
is at least the true count. Collisions only add. The minimum is therefore the
tightest upper bound available, and it is never an underestimate.

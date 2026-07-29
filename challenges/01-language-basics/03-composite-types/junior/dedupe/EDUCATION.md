# Deduplication

## Intuition

Remove duplicates by tracking what you've seen. A set (`map[T]struct{}`) gives
O(1) membership and preserves first-seen order when you append on first sight:

```go
if _, ok := seen[x]; !ok { out = append(out, x); seen[x] = struct{}{} }
```

## Approach

1. Keep a seen map[int]bool and an empty result slice.
2. Range the input; for each value not yet seen, mark it seen and append it.
3. Never write to the input, so it stays unmutated.
4. Return result.

## Solution

```go
func Dedupe(in []int) []int {
	result := []int{}
	seen := make(map[int]bool)
	for _, v := range in {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
```

## Walkthrough

Dedupe([5,4,5,4]): 5 new -> append(5); 4 new -> append(4); 5 seen -> skip; 4 seen -> skip. Result [5,4].

## Pitfalls

- Emit on **absence**, then record.
- Order-preserving dedup differs from sort-then-compact.
- Element type must be comparable to be a set key.

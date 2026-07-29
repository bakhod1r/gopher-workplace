# Prefix by predicate

## Intuition

Scan until the predicate fails, then slice up to that point:

```go
i := 0
for i < len(xs) && xs[i] > 0 { i++ }
return xs[:i]
```

## Approach

1. Start empty result.
2. Walk xs; break at the first element <=0.
3. Otherwise append it.
4. Return the leading positive run.

## Solution

```go
func TakePositive(xs []int) []int {
	out := []int{}
	for _, v := range xs {
		if v <= 0 {
			break
		}
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

[1,2,3,-1,4]: append 1,2,3; -1<=0 break -> [1,2,3].

## Pitfalls

- `xs[:i]` shares the backing array.
- Return a non-nil empty when the first element fails (`xs[:0]`).
- The dual, "drop while", slices from `i` onward.

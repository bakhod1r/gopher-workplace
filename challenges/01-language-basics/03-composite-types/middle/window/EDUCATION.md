# Rolling window sums

## Intuition

Naively each window sum is O(k); rolling makes it O(1) per step by adding the new
element and subtracting the one that left:

```go
sum := 0
for i := 0; i < k; i++ { sum += xs[i] }
out = append(out, sum)
for i := k; i < len(xs); i++ { sum += xs[i] - xs[i-k]; out = append(out, sum) }
```

## Approach

1. If k<=0 or k>len return empty.
2. Sum the first window.
3. Slide: add xs[i], subtract xs[i-k] (rolling sum).
4. Append each window sum.
5. Return sums.

## Solution

```go
func Sums(xs []int, k int) []int {
	out := []int{}
	if k <= 0 || k > len(xs) {
		return out
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += xs[i]
	}
	out = append(out, sum)
	for i := k; i < len(xs); i++ {
		sum += xs[i] - xs[i-k]
		out = append(out, sum)
	}
	return out
}
```

## Walkthrough

[1,2,3,4] k=2: first 1+2=3; +3-1=5; +4-2=7 -> [3,5,7].

## Pitfalls

- There are `len(xs)-k+1` windows; guard `k > len` and `k <= 0`.
- Rolling avoids recomputation but accumulates rounding for floats.
- Sub-slices would share memory; here you output computed sums.

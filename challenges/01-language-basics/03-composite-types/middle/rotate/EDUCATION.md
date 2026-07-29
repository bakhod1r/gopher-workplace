# Rotation by slice reassembly

## Intuition

Rotating left by `k` moves the first `k` elements to the end. Normalize `k` into
`[0,n)` first (handles large and negative), then concatenate the two parts:

```go
k = ((k % n) + n) % n
out := append([]int{}, xs[k:]...)
out = append(out, xs[:k]...)
```

## Approach

1. n=len; if 0 return empty.
2. Normalize k = ((k%n)+n)%n.
3. Result = xs[k:] followed by xs[:k], copied into a new slice.
4. Return it (input untouched).

## Solution

```go
func Left(xs []int, k int) []int {
	n := len(xs)
	if n == 0 {
		return []int{}
	}
	k = ((k % n) + n) % n
	out := make([]int, 0, n)
	out = append(out, xs[k:]...)
	out = append(out, xs[:k]...)
	return out
}
```

## Walkthrough

[1..5] k=2: xs[2:]=[3,4,5] + xs[:2]=[1,2] -> [3,4,5,1,2]. k=-1,n=3 -> ((-1%3)+3)%3=2 -> xs[2:]+xs[:2]=[3]+[1,2]=[3,1,2].

## Pitfalls

- Guard `n == 0` before `k % n`.
- `append([]int{}, xs[k:]...)` copies, so the result doesn't alias `xs`.
- An in-place rotate (three reversals) is O(1) space but mutates the input.

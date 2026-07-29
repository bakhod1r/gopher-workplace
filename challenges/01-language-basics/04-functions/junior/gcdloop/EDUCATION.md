# Euclidean algorithm

## Intuition

gcd(a,b) equals gcd(b, a mod b); iterating shrinks the pair until the remainder is 0.

## Approach

1. Loop while `b != 0`.
2. Replace `(a, b)` with `(b, a % b)`.
3. Return `a` when `b` reaches 0.

## Solution

```go
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## Walkthrough

`GCD(12, 8)`: (12,8)→(8,4)→(4,0), so the answer is 4.

## Pitfalls

- gcd(0, n) is n — the loop handles it since b becomes 0 immediately.
- Parallel assignment avoids a temporary.

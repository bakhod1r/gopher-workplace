# In-place transform via a callback

## Intuition

Reading `*p`, applying `f`, and storing back mutates the caller's variable with pluggable logic.

## Approach

1. Read `*p`.
2. Transform it with `f`.
3. Store the result back: `*p = f(*p)`.

## Solution

```go
func Apply(p *int, f func(int) int) {
	*p = f(*p)
}
```

## Walkthrough

`Apply(&x, square)` with `x = 3`: `f(3)` is `9`, written back into `x`.

## Pitfalls

- The pointee is both input and output.
- `f` must be non-nil; calling a nil func panics.

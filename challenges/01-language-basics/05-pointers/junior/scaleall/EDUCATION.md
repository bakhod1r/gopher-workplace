# Collections of pointers

## Intuition

A `[]*int` holds addresses; iterating and dereferencing mutates each target, with nil entries needing a guard.

## Approach

1. Range over the slice of pointers.
2. Skip nil entries with `if p != nil`.
3. Scale in place: `*p *= k`.

## Solution

```go
func ScaleAll(ps []*int, k int) {
	for _, p := range ps {
		if p != nil {
			*p *= k
		}
	}
}
```

## Walkthrough

For `[]*int{&a, &b}` with `k = 10`: first `*p *= 10` sets `a = 20`, then the second sets `b = 30`.

## Pitfalls

- Dereferencing a nil element panics — skip it.
- Mutating `*p` changes the original variable.

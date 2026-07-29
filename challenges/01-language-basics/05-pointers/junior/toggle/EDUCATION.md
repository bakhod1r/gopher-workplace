# Pointers to any type

## Intuition

A `*bool` mutates a caller's flag just as `*int` mutates an integer.

## Approach

1. `p` points at a bool.
2. `*p = !*p` flips it in place.

## Solution

```go
func Toggle(p *bool) {
	*p = !*p
}
```

## Walkthrough

`Toggle(&b)` with `b = false`: `!*p` is `true`, stored back into `b`.

## Pitfalls

- `!*p` reads the current bool then negates.
- Applies to every type, not just numbers.

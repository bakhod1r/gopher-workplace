# Pointers as mutable references

## Intuition

Passing `&x` lets a function modify the caller's variable through the pointer; a plain `int` parameter could not.

## Approach

1. `p` aliases the caller's int.
2. `*p++` increments it in place.

## Solution

```go
func Bump(p *int) {
	*p++
}
```

## Walkthrough

`Bump(&x)` with `x = 5` writes 6 back through the pointer.

## Pitfalls

- Forgetting the `*` writes to the pointer variable, not the pointee.
- A nil pointer dereference panics; callers must pass a valid address.

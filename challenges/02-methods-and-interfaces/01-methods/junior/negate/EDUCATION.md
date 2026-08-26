# Pointer Receiver on Defined Types

## Intuition

Even defined types like `type MyInt int` can have pointer receivers. `*MyInt`
gives the method access to the original value, allowing in-place mutation.

## Approach

1. Dereference `n`.
2. Negate it.
3. Store back.

## Solution

```go
func (n *MyInt) Negate() {
	*n = -(*n)
}
```

## Walkthrough

For `n := MyInt(5); n.Negate()`:
- `*n` is `5`.
- `-5` is stored back via `*n = -5`.
- `n` is now `-5`.

## Pitfalls

- Value receiver `(n MyInt)` — the negation is lost.
- Forgetting the dereference: `n = -n` on a pointer doesn't compile.
- `*n = -*n` is valid Go shorthand for `*n = -(*n)`.

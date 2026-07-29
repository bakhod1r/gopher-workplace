# Pointers as references

## Intuition

A pointer holds the address of a value; dereferencing (`*p`) reaches the value so a function can mutate the caller's variable.

## Approach

1. `p` has type `*int`: it points at the caller's variable.
2. Dereference with `*p` to reach the underlying int.
3. Add one **in place** so the store lands back at the same address.

## Solution

```go
func Inc(p *int) {
	*p++
}
```

## Walkthrough

Take `x := 41` and call `Inc(&x)`:

- `p` receives the **address** of `x`.
- `*p` reads the current value, `41`.
- `*p++` writes `42` back to that address.
- Control returns; in the caller, `x` is now `42`.

## Pitfalls

- `p++` moves the pointer variable; `*p++` changes the pointee.
- A nil pointer dereference panics.

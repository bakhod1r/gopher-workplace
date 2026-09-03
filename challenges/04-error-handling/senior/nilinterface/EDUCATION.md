# Typed Nil Trap

## Intuition

An interface value holds a type and a value. Assigning a nil `*OpError` fills in the type, so the interface is non-nil while the pointer inside is nil — and every `if err != nil` upstream fires.

## Approach

1. Check the pointer against nil explicitly.
2. Return the untyped nil in that case.
3. Return the pointer otherwise.

## Solution

```go
if e == nil {
	return nil
}
return e
```

## Walkthrough

`Wrap(e)` with a declared-but-unset `*OpError` returns the untyped nil, so the caller's comparison against nil succeeds.

## Pitfalls

- Returning `e` unconditionally, producing a non-nil error for a nil pointer.
- Declaring a function that returns `*OpError` and assigning it straight into an `error` variable at the call site.
- Comparing the interface against nil inside the method, where the receiver is already typed.

# Writing through a pointer

## Intuition

`*p = v` stores v at the address p holds, mutating the caller's variable.

## Approach

1. `p` points at the caller's int.
2. `*p = 0` stores zero at that address.
3. The change is visible to the caller because the storage is shared.

## Solution

```go
func Zero(p *int) {
	*p = 0
}
```

## Walkthrough

`x := 99`, `Zero(&x)`: `*p = 0` writes `0` at the address of `x`, so the caller reads `0` afterwards.

## Pitfalls

- Only `*p = ...` reaches the pointee.
- Passing the value by copy would not clear the original.

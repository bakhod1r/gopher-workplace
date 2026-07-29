# Short-circuit evaluation as a guard

## Intuition

`&&` and `||` evaluate left to right and stop early; the nil/bounds guard must precede the operation it protects.

## Approach

1. `&&` short-circuits left to right; the nil check must come **first**.
2. The bug writes `*p > 0 && p != nil`, dereferencing before checking nil.
3. Reorder to `p != nil && *p > 0`.

## Solution

```go
func ValueOr(p *int, def int) int {
	if p != nil && *p > 0 {
		return *p
	}
	return def
}
```

## Walkthrough

`ValueOr(nil, 5)` dereferences nil in the bug's first operand and panics. Checking `p != nil` first stops evaluation before the deref.

## Pitfalls

- `p != nil && *p > 0` is safe; the reverse panics on nil.
- Rely on left-to-right short-circuit ordering deliberately.

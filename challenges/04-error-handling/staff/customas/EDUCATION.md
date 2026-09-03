# Custom As Method

## Intuition

`As` lets an error present itself as a type it is not, which is how two generations of an API coexist. The method is consulted before the ordinary type match.

## Approach

1. Assert `target` to `**Modern`.
2. Return false when it is anything else.
3. Assign a converted value through the pointer and return true.

## Solution

```go
m, ok := target.(**Modern)
if !ok {
	return false
}
*m = &Modern{Op: e.Op, Code: e.Num}
return true
```

## Walkthrough

Asking for the concrete `*LegacyError` still works: `errors.As` falls back to the ordinary type match when the custom method declines.

## Pitfalls

- Asserting to `*Modern` instead of `**Modern`.
- Returning true without assigning, leaving the caller with a nil pointer.
- Matching every target type, hijacking unrelated lookups.

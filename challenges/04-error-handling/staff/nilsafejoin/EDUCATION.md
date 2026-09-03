# Join With Typed Nils

## Intuition

A typed nil is a non-nil interface value, so every generic filter that checks `err != nil` lets it through. The only reliable defence is to know the concrete types you accept and check them.

## Approach

1. Skip untyped nil entries.
2. Skip entries that assert to a nil `*OpError`.
3. Join what remains.

## Solution

```go
var kept []error
for _, err := range errs {
	if err == nil {
		continue
	}
	if oe, ok := err.(*OpError); ok && oe == nil {
		continue
	}
	kept = append(kept, err)
}
return errors.Join(kept...)
```

## Walkthrough

The typed nil passes `err != nil` but fails the second check, so the join receives only the real failure and its message is a single line.

## Pitfalls

- Trusting `errors.Join` to drop typed nils.
- Calling `Error()` on the entry to test it, which panics for a nil receiver that dereferences a field.
- Filtering at every call site instead of once at the boundary.

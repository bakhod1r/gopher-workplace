# Email Shape

## Intuition

Splitting a string and validating its pieces is two questions: was there a separator at all, and are the pieces meaningful? Distinct errors let the caller word the message precisely.

## Approach

1. Cut on `"@"`.
2. Return `ErrNoAt` when the separator was not found.
3. Return `ErrEmptyPart` when either side is empty.
4. Return nil otherwise.

## Solution

```go
local, domain, found := strings.Cut(s, "@")
if !found {
	return ErrNoAt
}
if local == "" || domain == "" {
	return ErrEmptyPart
}
return nil
```

## Walkthrough

For `"@"` the cut succeeds with both parts empty, so the second guard fires with `ErrEmptyPart`.

## Pitfalls

- Using `strings.Split` and forgetting the no-separator case yields one element.
- Returning `ErrNoAt` for `"@b.com"` — the `@` is present, the local part is not.
- Attempting full RFC validation; the contract is a shape check.

# Key And Value

## Intuition

Cutting at the *first* separator is what makes `URL=a=b` parse correctly: everything after the first `=` belongs to the value, separators included.

## Approach

1. Cut on `"="`.
2. Return `ErrNoSeparator` when not found.
3. Return `ErrEmptyKey` when the key is empty.
4. Return key, value, nil.

## Solution

```go
k, v, found := strings.Cut(s, "=")
if !found {
	return "", "", ErrNoSeparator
}
if k == "" {
	return "", "", ErrEmptyKey
}
return k, v, nil
```

## Walkthrough

`"URL=a=b"` cuts into `"URL"` and `"a=b"`; the second `=` stays inside the value.

## Pitfalls

- Using `strings.Split`, which shreds `a=b` into extra pieces.
- Rejecting an empty value — `KEY=` is legal.
- Returning the partial key alongside an error.

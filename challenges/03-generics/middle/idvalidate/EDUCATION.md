# String-Like Identifiers

## Intuition

Type parameters do not automatically satisfy concrete parameter types, so the conversion at the boundary is required — and it is free.

## Approach

1. Convert `v` to `string`.
2. Return non-empty and prefixed.

## Solution

```go
func ValidID[T ~string](v T, prefix string) bool {
	s := string(v)
	return s != "" && strings.HasPrefix(s, prefix)
}
```

## Walkthrough

`ValidID(UserID("u_1"), "u_")` converts to `"u_1"`, which is non-empty and prefixed.

## Pitfalls

- Passing `v` straight to `strings.HasPrefix`, which does not compile.
- Using `string` instead of `~string`, rejecting every named ID type.
- Treating an empty ID as valid when it happens to have an empty prefix.

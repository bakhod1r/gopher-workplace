# Dereference Or Default

## Intuition

A pointer type parameter behaves like any pointer: comparable to `nil`, dereferenced with `*`. The generic part adds no new rules.

## Approach

1. Return `def` when `p == nil`.
2. Otherwise return `*p`.

## Solution

```go
func Deref[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

func Ptr[T any](v T) *T {
	return &v
}
```

## Walkthrough

`Deref((*int)(nil), 0)` takes the guard branch and returns `0` instead of panicking with a nil dereference.

## Pitfalls

- Dereferencing before the nil check.
- Comparing `*p` to a zero value to detect nil — that already panicked.
- Returning `*p` when `p` is nil because the pointer "looked" set.

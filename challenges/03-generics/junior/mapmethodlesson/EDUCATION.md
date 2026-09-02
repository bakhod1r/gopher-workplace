# Map Needs A Function

## Intuition

Method sets have to be known when the type is instantiated, so Go forbids per-method type parameters. Anything needing a fresh type parameter must be a free function.

## Approach

1. `MapSlice`: allocate a `Slice[U]`, append `f(v)` per element.
2. `Each`: range and call `f`.

## Solution

```go
func MapSlice[T, U any](s Slice[T], f func(T) U) Slice[U] {
	out := make(Slice[U], 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func (s Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
```

## Walkthrough

`MapSlice(Slice[int]{1,2}, itoa)` returns `Slice[string]{"1","2"}` — a different instantiation, which no method could have named.

## Pitfalls

- Trying `func (s Slice[T]) Map[U any](f func(T) U) Slice[U]` — a compile error.
- Returning `[]U` instead of `Slice[U]` and losing the methods.
- Making `Each` return a value nobody asked for.

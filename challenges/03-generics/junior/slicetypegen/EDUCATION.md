# Method On A Slice Type

## Intuition

Attaching methods requires a named type, and generics let that named type stay parameterised. `Slice[int]` and `[]int` are convertible, so callers lose nothing.

## Approach

1. `Filter`: allocate a `Slice[T]`, append what `keep` accepts, return it.
2. `Len`: return `len(s)`.

## Solution

```go
func (s Slice[T]) Filter(keep func(T) bool) Slice[T] {
	out := make(Slice[T], 0, len(s))
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func (s Slice[T]) Len() int {
	return len(s)
}
```

## Walkthrough

`Slice[int]{1,2,3}.Filter(isEven)` returns `Slice[int]{2}`, which still has `Len` available.

## Pitfalls

- Returning `[]T`, which breaks chaining.
- Appending into the receiver, which can overwrite the caller's backing array.
- Trying to declare `func (s Slice[T]) Map[U any](...)` — methods take no new type parameters.

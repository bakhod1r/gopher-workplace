# Reactive Stream

## Intuition

Chaining is not a language feature — it is a convention: every operator returns
the same pointer it was called on, so the expression's type never changes. Once
that holds, arbitrary pipelines compose.

The in-place filter is the interesting mechanic. Because `append` to a
zero-length reslice writes into the array you are already reading, and the write
cursor can never get ahead of the read cursor, the whole filter runs with no
allocation at all.

## Approach

1. `Filter`: reslice to zero length, append the survivors, store the result back.
2. `Map`: walk by index and overwrite each slot.
3. Both return `s`.

## Solution

```go
func (s *Stream) Filter(fn func(int) bool) *Stream {
	out := s.Data[:0]
	for _, v := range s.Data {
		if fn(v) {
			out = append(out, v)
		}
	}
	s.Data = out
	return s
}

func (s *Stream) Map(fn func(int) int) *Stream {
	for i := range s.Data {
		s.Data[i] = fn(s.Data[i])
	}
	return s
}
```

## Walkthrough

`[1 2 3 4]` filtered by "even": the loop reads index 0 (1, dropped), index 1
(2, appended at write position 0), index 2 (3, dropped), index 3 (4, appended at
write position 1). The array is now `[2 4 3 4]` and `out` has length 2, so
`s.Data` becomes `[2 4]`.

`Map` then multiplies both survivors in place: `[20 40]`. `reflect.DeepEqual`
compares length and elements, so the stale `3, 4` beyond the length are
invisible.

## Pitfalls

- **Forgetting `return s`.** The chain does not compile, or returns nil.
- **`out := s.Data[:0]` but never assigning `s.Data = out`.** The survivors are
  written into the array, but the length never shrinks: `[2 4 3 4]`.
- **`for _, v := range` in `Map`.** Mutates copies; the data is unchanged.
- **Assuming the old elements are gone.** They are still in the backing array
  beyond `len`. If they were pointers, that would keep objects alive — a real
  leak. Zero the tail when the elements hold references.

## Eager, not reactive

A true reactive stream is lazy and push-based: operators are composed first and
values flow through them one at a time. This one is eager, so each operator does
a full pass. The chaining shape is the same, which is what makes the eager
version a good place to learn it.

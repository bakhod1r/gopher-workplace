# Enumerable Interface

## Intuition

An interface can require several methods. `Join` needs both a size and an accessor, so the contract lists both.

## Approach

1. `Len` returns `len(w)`; `At` returns `w[i]`.
2. In `Join`, loop `i` from 0 to `e.Len()-1`.
3. Prepend `sep` when `i > 0`, then append `e.At(i)`.

## Solution

```go
func (w Words) Len() int { return len(w) }

func (w Words) At(i int) string { return w[i] }

func Join(e Enumerable, sep string) string {
	out := ""
	for i := 0; i < e.Len(); i++ {
		if i > 0 {
			out += sep
		}
		out += e.At(i)
	}
	return out
}
```

## Walkthrough

For `Words{"a","b","c"}` with `sep = "-"`: i=0 adds `a`; i=1 adds `-` then `b`; i=2 adds `-` then `c` — `"a-b-c"`.

## Pitfalls

- Appending `sep` after every item leaves a trailing separator.
- Calling `e.Len()` is fine in the loop condition here, but caching it is clearer for large inputs.
- Implementing only `Len` — the type then fails to satisfy `Enumerable`.

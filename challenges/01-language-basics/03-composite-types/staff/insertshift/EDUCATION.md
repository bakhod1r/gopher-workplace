# Opening a gap for insertion

## The idea

In-place insertion at `i` first grows the slice by one, then shifts the tail one
step **right** to open a gap:

```go
xs = append(xs, 0)
copy(xs[i+1:], xs[i:]) // dst starts one past src -> shift right
xs[i] = v
```

`copy` is overlap-safe (memmove semantics), so the right-shift doesn't corrupt.

## Why it matters

This is what `slices.Insert` does. The direction (`dst = xs[i+1:]`, `src =
xs[i:]`) is the crux; reversing it shifts left and loses data.

## Watch out

- Insert shifts right; delete shifts left.
- Grow the slice before shifting, or you write past the end.
- `copy` copies `min(len(dst),len(src))` — the counts line up here.

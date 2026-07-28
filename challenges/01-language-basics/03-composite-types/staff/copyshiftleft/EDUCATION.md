# Overlapping copy and direction

## The idea

`copy(dst, src)` copies `min(len)` elements and handles overlap like `memmove`.
To shift left (drop index 0), the destination is the earlier region:

```go
copy(xs, xs[1:]) // xs[i] = xs[i+1]
xs[len(xs)-1] = 0
```

`copy(xs[1:], xs)` shifts the other way, replicating `xs[0]`.

## Why it matters

In-place shifts (ring buffers, sliding windows, deletion) hinge on getting the
copy direction right. `copy` is overlap-safe, so the only bug is swapping src/dst.

## Watch out

- Left shift: `copy(xs, xs[1:])`. Right shift: `copy(xs[1:], xs)`.
- `copy` returns the count copied (`len-1` here).
- Remember to clear the vacated slot to avoid a stale/duplicated value.

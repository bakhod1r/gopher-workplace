# Mapped Region Reads

## Intuition

`ReadAt` is `Read` without a cursor: every call is absolute, so it is safe to
call concurrently and there is no state to get wrong. What people *do* get wrong
is the end-of-region contract, because it is the rare Go API where a non-nil
error accompanies real, usable data.

## Approach

1. Reject offsets outside the region before any slicing.
2. `copy` from `m.Data[off:]` into `p` — it naturally truncates.
3. Report `io.EOF` when fewer bytes were delivered than requested.

## Solution

```go
func (m *Mmap) ReadAt(p []byte, off int) (int, error) {
	if off < 0 || off >= len(m.Data) {
		return 0, io.EOF
	}
	n := copy(p, m.Data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}
```

## Walkthrough

- `Data` of 5, `p` of 3, `off` 1: `m.Data[1:]` has 4 bytes, `p` has 3, so `copy`
  moves 3 and `n == len(p)` — a full read, no error.
- `Data` of 3, `p` of 4, `off` 1: the source has 2 bytes left, so `copy` moves 2.
  `n < len(p)`, so the caller gets `(2, io.EOF)` — the two bytes are valid and
  the error explains why there were not four.
- `off == 3` on a 3-byte region: the guard fires before `m.Data[3:]`, which
  would itself be legal (an empty slice) but would then produce `(0, nil)` — an
  infinite loop for a caller that reads until EOF.

## Pitfalls

- **Slicing before the bounds check.** `m.Data[-1:]` panics; `m.Data[99:]` panics.
- **Returning `(0, io.EOF)` on a short read.** Throws away bytes that were
  actually copied; a caller looping on `ReadAt` silently loses data.
- **Returning `(n, nil)` at the exact end.** Callers that stop only on EOF spin
  forever.
- **Treating `err != nil` as "ignore n".** For `ReaderAt` and `Reader` alike, Go
  requires processing `n` bytes *before* considering the error.

## Why this is the `io.ReaderAt` shape

`io.ReaderAt` uses `int64` offsets; this puzzle uses `int` to stay inside the
topic. The semantics are otherwise those the standard library documents — and
`os.File`, `bytes.Reader` and `golang.org/x/exp/mmap` all implement exactly this
contract, which is why matching it makes a type drop-in compatible with
`io.SectionReader` and friends.

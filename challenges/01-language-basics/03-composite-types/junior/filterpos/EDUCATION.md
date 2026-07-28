# Building slices with append

## The idea

`append` adds to a slice and returns a (possibly relocated) new header — always
assign the result back:

```go
out := []int{}
for _, x := range xs {
	if x > 0 { out = append(out, x) }
}
```

## Why it matters

Filter/map/collect all build a result incrementally with `append`. Forgetting to
capture `append`'s return, or starting from `nil` when the caller distinguishes
nil from empty, are common slice bugs.

## Watch out

- `append(s, x)` may reallocate; the returned slice is authoritative.
- `[]int{}` is non-nil length 0; `var out []int` is nil — pick per the test/spec.
- Pre-sizing with `make([]int, 0, len(xs))` avoids regrowth in hot paths.

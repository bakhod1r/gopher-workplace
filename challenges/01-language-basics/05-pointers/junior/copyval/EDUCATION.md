# Copying values via pointers

## Intuition

`*dst = *src` reads the source pointee and writes it to the destination pointee — a value copy, not an alias.

## Approach

1. Read `*src`.
2. Write it into `*dst`.
3. `src` is untouched.

## Solution

```go
func CopyInto(dst, src *int) {
	*dst = *src
}
```

## Walkthrough

`CopyInto(&a, &b)` with `a = 1`, `b = 9`: `*dst = *src` copies `9` into `a`; `b` stays `9`.

## Pitfalls

- After the copy the two remain separate ints.
- Order matters: dst is the target.

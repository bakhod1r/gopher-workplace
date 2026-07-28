# Slice Copy Independence

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 12 min

## Context

A cache hands out the first N items of an internal slice. Returning `s[:n]`
looked fine until callers mutated the result and silently corrupted the cache —
a sub-slice shares the same backing array. The result must be independent.

## Task

Implement `Head` in [subslice.go](subslice.go) so it returns the first `n`
elements of `s` as an **independent copy**: writes to the result must not touch
`s`, and vice-versa. Clamp `n` to `len(s)`; treat `n <= 0` as empty.

Do **not** change the function signature or the tests.

## Examples

```go
Head([]int{1, 2, 3, 4}, 2) // => []int{1, 2}
Head([]int{1, 2}, 5)       // => []int{1, 2}   (clamped)
Head([]int{1, 2, 3}, 0)    // => []int{}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice expressions** | `s[:n]` makes a new header over the *same* backing array — no data is copied. |
| 2 | **Shared backing = shared writes** | Mutating a sub-slice element is visible through the original, and vice-versa. |
| 3 | **`make` + `copy`** | `out := make([]int, n); copy(out, s)` gives a separate backing array. |
| 4 | **Bounds & clamping** | `s[:n]` panics when `n > len(s)`; clamp `n` first. |

## Hint

Clamp `n` into `[0, len(s)]`, allocate `make([]int, n)`, then `copy` from `s`.
Returning `s[:n]` shares the backing array and fails the independence test.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

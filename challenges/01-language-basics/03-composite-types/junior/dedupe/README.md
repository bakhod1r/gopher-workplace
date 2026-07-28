# Slice Dedupe

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 15 min

## Context

A data-ingestion step must drop duplicate integers from a slice while keeping
the order in which each value first appeared — and it must leave the caller's
slice untouched.

## Task

Implement `Dedupe` in [dedupe.go](dedupe.go) so that it:

1. Returns elements in **first-appearance order**, duplicates removed.
2. **Does not mutate** the input slice.
3. Returns an empty (non-nil) slice for `nil` or empty input.

Do **not** change the function signature or the tests.

## Examples

```go
Dedupe([]int{1, 1, 2, 3, 3, 3}) // => []int{1, 2, 3}
Dedupe([]int{5, 4, 5, 4})       // => []int{5, 4}
Dedupe([]int{7, 7, 7})          // => []int{7}
Dedupe(nil)                     // => []int{}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice header** | A slice is a pointer+len+cap view over a backing array, not the array itself. |
| 2 | **Copy vs alias** | `out := in` copies only the header — both still share one backing array. |
| 3 | **`append` semantics** | `append` may reuse spare capacity, silently overwriting data another slice sees; start from a fresh slice to stay safe. |
| 4 | **Map as a set** | `map[int]bool` (or `map[int]struct{}`) tracks seen values in O(1). |
| 5 | **nil vs empty** | Both have `len == 0`; return `[]int{}`, not `nil`, for the empty case. |

## Hint

Build a *new* result slice (`make([]int, 0, len(in))` or `[]int{}`) and a
`seen` map. Range over the input, appending a value the first time you see it.
Never append into a slice that aliases the input.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

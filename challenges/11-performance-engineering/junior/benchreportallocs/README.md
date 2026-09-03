# What `-benchmem` Prints

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`go test -bench . -benchmem` appends two columns to every benchmark line: bytes allocated per operation and allocation count per operation. Both are totals divided by the iteration count, truncated. Reproduce that formatting.

## Task

Implement `Report` in [benchreportallocs.go](benchreportallocs.go):

1. Divide `bytes` and `allocs` by `iters`, truncating toward zero.
2. Return `"<bytes> B/op\t<allocs> allocs/op"` — a single tab between the two fields.
3. A non-positive `iters` reports `"0 B/op\t0 allocs/op"`.

## Examples

**Example 1:**

```
Input:  Report(2048, 8, 4)
Output: "512 B/op\t2 allocs/op"
```

**Example 2:**

```
Input:  Report(10, 3, 4)
Output: "2 B/op\t0 allocs/op"
```

**Example 3:**

```
Input:  Report(999, 999, 0)
Output: "0 B/op\t0 allocs/op"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Per-op is a mean** | Both columns are cumulative totals divided by `b.N`. |
| 2 | **Truncation hides small allocations** | `3 allocs / 4 iters` prints `0 allocs/op`, not `1`. |
| 3 | **Divide-by-zero guard** | A benchmark that never ran must report zeros, not panic. |

## Topics used again

`fmt.Sprintf`, integer division.

## Hint

Guard `iters` first, then one `fmt.Sprintf` with a `\t` in the middle.

## Validate

```bash
make verify
```

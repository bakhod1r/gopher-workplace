# Storage Reporter

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A storage report is built section by section, and each section lists the byte
sizes of its files. The report converts every file to whole kilobytes and totals
each section. Sections are processed concurrently, and inside a section the
individual file conversions run concurrently too.

## Task

Implement `SectionKilobytes` in [storagereporter.go](storagereporter.go) so that:

1. Return a slice with one total per section, in section order.
2. The total of a section is the sum of `bytes / 1024` for each of its files; an empty section totals `0`.
3. Use one goroutine per section and, inside it, one goroutine per file — each level with its own `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SectionKilobytes([][]int{{2048, 1024}, {4096}})
Output: [3 4]
```

**Example 2:**

```
Input:  SectionKilobytes([][]int{{}})
Output: [0]
```

**Example 3:**

```
Input:  SectionKilobytes(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Nested WaitGroups** | The inner group must be waited on *inside* the outer goroutine, before it writes its total. |

## Hint

Each outer goroutine owns a private `kb` slice and a private `inner` WaitGroup.
Call `inner.Wait()` before summing, or you will sum zeros.

## Validate

```bash
make verify
```

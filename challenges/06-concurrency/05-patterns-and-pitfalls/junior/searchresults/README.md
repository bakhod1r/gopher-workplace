# Top Search Results

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The search API renders one page of results, but the index goroutine keeps
streaming every match it finds. If the handler simply stops reading after ten
hits, that producer stays blocked on its next send for the life of the
process — a goroutine leak that shows up as slowly growing memory.

## Task

Implement `TopResults` in [searchresults.go](searchresults.go) so that:

1. For `n <= 0` it takes nothing.
2. Otherwise it appends hits until it has `n` of them, then stops taking.
3. In every case it drains the remaining hits so the producing goroutine can run to completion and exit.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  stream of a, b, c, d, e with n = 3
Output: []string{"a", "b", "c"}
```

**Example 2:**

```
Input:  stream of a, b with n = 9
Output: []string{"a", "b"}
```

**Example 3:**

```
Input:  stream of a, b with n = 0
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Take-n** | Consuming a bounded prefix of an unbounded stream. |
| 2 | **Goroutine leak** | An abandoned producer blocks on send forever; draining releases it. |
| 3 | **Empty range** | `for range hits {}` is the idiomatic drain-and-discard loop. |

## Hint

Two loops: the first breaks once you hold `n` hits, the second is an empty
`for range hits {}` that lets the producer finish.

## Validate

```bash
make verify
```

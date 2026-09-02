# Email Normalizer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A CRM import cleans a batch of contact addresses before de-duplication:
surrounding whitespace is stripped and the address is lowercased so that
`A@X.io` and `a@x.io ` collapse to the same key. Addresses are cleaned
concurrently and the batch keeps its import order.

## Task

Implement `Normalize` in [emailnormalizer.go](emailnormalizer.go) so that:

1. Return a new slice the same length as `addrs`; do not modify the input.
2. Address `i` is `addrs[i]` with surrounding whitespace removed and then lowercased.
3. Normalise each address in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Normalize([]string{" A@X.io "})
Output: [a@x.io]
```

**Example 2:**

```
Input:  Normalize([]string{"   "})
Output: []
```

**Example 3:**

```
Input:  Normalize(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Composing pure calls** | `TrimSpace` then `ToLower` — both allocate new strings, so neither can be observed half-done. |

## Hint

`strings.TrimSpace` already handles tabs and newlines; you do not need to list
whitespace characters yourself.

## Validate

```bash
make verify
```

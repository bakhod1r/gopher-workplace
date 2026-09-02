# Collect Provider Failures

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

A checkout splits a payment across several providers and calls them in
parallel. One declined card must not hide another: the operator needs *every*
failure in the reconciliation report, in a stable order they can diff against
yesterday's run.

## Task

Implement `ChargeAll` in [paymentcharge.go](paymentcharge.go) so that:

1. It charges each provider in its own goroutine, tracked by a `WaitGroup`.
2. A goroutine whose charge returns nil contributes nothing.
3. Failure messages are appended under a mutex and the collected slice is sorted before it is returned.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ChargeAll([]string{"ok-visa", "bad-amex"}, charge)
Output: []string{"bad-amex declined"}
```

**Example 2:**

```
Input:  ChargeAll([]string{"bad-z", "bad-a"}, charge)
Output: []string{"bad-a declined", "bad-z declined"}
```

**Example 3:**

```
Input:  ChargeAll(nil, charge)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Collecting from many goroutines** | `append` to a shared slice needs a mutex — the slice header is shared state. |
| 2 | **Sorted output** | Failures complete in arbitrary order, so sort before returning. |
| 3 | **Aggregate errors** | Report every failure; the first one is rarely the whole story. |

## Hint

Call `charge` outside the lock, and take the mutex only around the `append`.
Sort once, after `wg.Wait()`.

## Validate

```bash
make verify
```

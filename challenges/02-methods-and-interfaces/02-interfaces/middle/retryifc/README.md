# Retry Interface

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A flaky operation is retried a bounded number of times before the caller gives up.

## Task

Implement the stub(s) in [retryifc.go](retryifc.go):

1. Implement `Do` on `*Flaky` — fail with `ErrTemporary` for the first `FailTimes` calls, then succeed with `Value`.
2. Implement `Retry`, which calls the operation up to `attempts` times and returns the first success.
3. Stop immediately on a non-temporary error; return the last error when every attempt fails.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Retry(&Flaky{FailTimes: 2, Value: "ok"}, 3)
Output: "ok", nil, 3 calls
```

**Example 2:**

```
Input:  Retry(&Flaky{FailTimes: 5}, 2)
Output: "", ErrTemporary, 2 calls
```

**Example 3:**

```
Input:  Retry(&Permanent{}, 5)
Output: "", ErrFatal, 1 call
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Operation as an interface** | Retry logic is written once for any operation. |
| 2 | **errors.Is for classification** | Reused: temporary versus fatal decides whether to retry. |
| 3 | **Bounded loops** | The attempt budget must be respected exactly. |

## Hint

A non-temporary error must return immediately — do not burn the remaining attempts.

## Validate

```bash
make verify
```

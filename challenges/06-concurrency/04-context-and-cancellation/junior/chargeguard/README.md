# Refuse to Charge a Dead Request

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The checkout service captures a card charge at the end of the request. If the client disconnected or the request budget expired while the fraud check was running, capturing anyway bills a customer who will never see a confirmation — and produces a support ticket and a refund. The guard has to run before the side effect.

## Task

Implement the exported function(s) in [chargeguard.go](chargeguard.go) so that:

1. It checks `ctx.Err()` before doing anything else.
2. If that error is non-nil it returns `"", err` and never calls `capture`.
3. Otherwise it returns `capture(), nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Charge(live ctx, capture)
Output: "captured", nil
```

**Example 2:**

```
Input:  Charge(cancelled ctx, capture)
Output: "", context.Canceled (capture not called)
```

**Example 3:**

```
Input:  Charge(expired ctx, capture)
Output: "", context.DeadlineExceeded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`ctx.Err()` as a cheap guard** | A non-blocking check for "is this work still wanted?". |
| 2 | **Guard before side effects** | Check first, act second — the ordering is the whole point. |
| 3 | **Propagating the reason** | Return `ctx.Err()` itself so the caller can classify it. |

## Hint

`if err := ctx.Err(); err != nil { return "", err }` and only then call `capture()`.

## Validate

```bash
make verify
```

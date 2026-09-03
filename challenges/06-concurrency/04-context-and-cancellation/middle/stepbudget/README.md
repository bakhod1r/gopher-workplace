# Sequential Step Budget

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

An order checkout runs three upstream calls in sequence — reserve stock, authorise payment, write the ledger — under one overall request deadline. Each step gets its own sub-context so a hung step cannot eat the whole budget, and the chain stops the moment the request's deadline is gone.

## Task

Implement the stubbed functions in [stepbudget.go](stepbudget.go) so that:

1. Check the parent context **before** each step; a finished parent returns `ctx.Err()` and stops.
2. Give every step its own cancellable child context, cancelled as soon as the step returns.
3. Stop at the first error and return it, along with how many steps ran.
4. Return `nil` when every step succeeded.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RunSteps(live ctx, [ok, ok])
Output: 2, nil
```

**Example 2:**

```
Input:  RunSteps(live ctx, [ok, bad, ok])
Output: 2, the bad step's error
```

**Example 3:**

```
Input:  RunSteps(cancelled ctx, [ok])
Output: 0, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Per-step sub-context | `WithCancel(ctx)` scopes one step's goroutines and cleanup; cancelling it does not touch the parent. |
| 2 | `cancel()` on every path | Cancel right after the step returns — a deferred cancel inside a loop piles up until the function ends. |
| 3 | Deadline inheritance | A child can only be shorter-lived than its parent, never longer. |
| 4 | Check before work | `ctx.Err()` at the top of each iteration turns a cancelled request into zero wasted upstream calls. |

## Hint

Loop: `if err := ctx.Err(); err != nil { return ran, err }`, then `stepCtx, cancel := context.WithCancel(ctx)`, run the step, call `cancel()` immediately — not with `defer`.

## Validate

```bash
make verify
```

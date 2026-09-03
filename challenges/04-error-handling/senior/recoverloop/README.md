# Isolate Each Item

**Level:** senior
**Topic:** 04-error-handling

## Context

A batch processor runs a handler per record. One record that panics must fail alone; the rest of the batch still runs.

## Task

Implement `Process` in [recoverloop.go](recoverloop.go):

1. Call `h` for every item, in order.
2. Convert a panic on one item into an error wrapping `ErrPanic`, then continue.
3. Return all failures joined, or nil when every item succeeded.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Process([]int{1, 2}, okH)
Output: nil
```

**Example 2:**

```
Input:  Process([]int{1, 2, 3}, panicOn2)
Output: 1 failure, 3 items handled
```

**Example 3:**

```
Input:  Process(nil, h)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Per-item recovery** | The recovery scope decides the blast radius. |
| 2 | **Closure per iteration** | A defer inside a helper runs per item. |
| 3 | **Collecting failures** | Isolation without silence. |

## Hint

A `defer` written directly in the loop body only runs when the whole function returns — the recovery needs its own function call per item.

## Validate

```bash
make verify
```

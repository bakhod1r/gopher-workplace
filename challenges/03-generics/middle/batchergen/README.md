# Batcher

**Level:** middle  
**Topic:** 03-generics

## Context

A writer sends records to a remote API in batches of 100, and the final partial batch must still be delivered at shutdown.

## Task

Implement the stub(s) in [batchergen.go](batchergen.go):

1. Implement `NewBatcher`, `Add`, and `Flush`.
2. `Add` returns a batch only when it becomes full; `Flush` returns any remainder.
3. A size below 1 is treated as 1.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  size 2: Add(1)
Output: nil, false
```

**Example 2:**

```
Input:  size 2: Add(1); Add(2)
Output: [1 2], true
```

**Example 3:**

```
Input:  Add(1); Flush()
Output: [1], true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Handing over ownership** | Returning the buffer and setting it to nil avoids a copy and prevents aliasing. |
| 2 | **Partial batches** | Without `Flush`, the tail of a stream would never be sent. |
| 3 | **Guarding the size** | A zero size would emit a batch on every add — or never. |

## Hint

Hand the buffer out and set the field to `nil` — the next `Add` allocates a fresh one.

## Validate

```bash
make verify
```

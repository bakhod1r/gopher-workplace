# Command Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A command is an action captured as a value, so it can be queued now and run
later. The `Invoker` collects commands and, on demand, fires them in order and
empties its queue.

## Task

Implement `ExecuteAll` on `*Invoker` in [commandpatt.go](commandpatt.go):

1. Call every command in `inv.commands`, in order.
2. Set `inv.commands` to `nil` afterwards.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(x += 5), Add(x *= 2), ExecuteAll()   // x starts at 0
Output: x == 10
```

**Example 2:**

```
Input:  ExecuteAll() on an empty invoker
Output: no-op, queue stays empty
```

**Example 3:**

```
Input:  ExecuteAll() twice in a row
Output: commands run once; the second call does nothing
```

_Explanation:_ clearing the queue is what makes the second call harmless.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Function type as value** | `Command` is `func()`; a slice of them is a queue of pending work. |
| 2 | **Order matters** | `x += 5` then `x *= 2` gives 10; the other order gives 5. |
| 3 | **Pointer receiver** | Clearing `inv.commands` must be visible to the caller. |

## Hint

`for _, c := range inv.commands { c() }`, then `inv.commands = nil`. Clearing
before running loses the queue.

## Validate

```bash
make verify
```

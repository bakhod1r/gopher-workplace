# Lock-Free Stack

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A lock-free push is a bet: read the current head, build a node pointing at it,
and swap it in — but only if nobody changed the head in the meantime. If someone
did, the bet is off and you retry. No goroutine ever blocks another.

## Task

Implement `Push` on `*Stack` in [lockfreestk.go](lockfreestk.go):

1. Loop.
2. Load the current head into `old` and set `n.next = old`.
3. `CompareAndSwap(old, n)`; on success, return. On failure, loop again.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(1) on an empty stack
Output: head is the new node, next == nil
```

**Example 2:**

```
Input:  Push(1); Push(2)
Output: head is 2, head.next is 1  (LIFO)
```

**Example 3:**

```
Input:  100 goroutines each Push once
Output: exactly 100 nodes on the list, no lost updates
```

_Explanation:_ every failed CAS is retried, so no push is dropped.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Compare-and-swap** | The atomic "swap only if unchanged" that the whole algorithm rests on. |
| 2 | **Re-reading inside the loop** | `old` must be reloaded on every attempt, or the retry uses a stale value forever. |
| 3 | **`atomic.Pointer[T]`** | Typed atomic pointer, no `unsafe` required. |

## Hint

Both the `Load` and the `n.next` assignment belong **inside** the loop. Hoisting
either one out turns the retry into an infinite loop that can never succeed.

## Validate

```bash
make verify
```

# Work-Stealing Deque

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Worker queues went idle while one worker had a backlog. Idle workers now steal from the far end of a busy worker's deque.

## Task

Implement the stub(s) in [workstealq.go](workstealq.go):

1. Implement `Push`, `Pop`, and `Steal` on `*Deque`: the owner pushes and pops at the bottom; thieves steal from the top.
2. `Pop` and `Steal` must never return the same item twice, including when they race over the last element.
3. Constraint: `-race` clean, no item lost or duplicated, and `Steal` must not block the owner's fast path.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(1), Push(2), Pop()
Output: 2 (LIFO for the owner)
```

**Example 2:**

```
Input:  Push(1), Push(2), Steal()
Output: 1 (FIFO for thieves)
```

**Example 3:**

```
Input:  Pop and Steal racing on one item
Output: exactly one of them wins
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Chase-Lev deque** | Owner-local LIFO, thief-facing FIFO, with a CAS on the last element. |
| 2 | **Contended last item** | The owner and a thief can both target one element; a CAS decides. |
| 3 | **Memory ordering** | Reused: atomics establish the ordering the algorithm depends on. |

## Hint

The owner decrements `bottom` first, then checks against `top`; if only one item remains, both sides race with a CAS on `top`.

## Validate

```bash
make verify
```

# Cyclic Barrier

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A simulation advances in rounds. Every worker must finish round N before any starts round N+1, and the barrier is reused every round.

## Task

Implement the stub(s) in [cyclicbarrier.go](cyclicbarrier.go):

1. Implement `Await` on `*Barrier`, blocking until `Parties` goroutines have arrived, then releasing them all and rearming for the next round.
2. `Await` returns the round index it completed.
3. Constraint: `-race` clean, correct across many rounds, and no goroutine may run ahead into the next round before the last arrival.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  3 parties, 3 Await calls
Output: all three are released
```

**Example 2:**

```
Input:  a second round on the same barrier
Output: works again, round index 1
```

**Example 3:**

```
Input:  round indexes returned
Output: identical for everyone in that round
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cyclic barrier** | A reusable rendezvous; the last arrival releases the group. |
| 2 | **Generation counters** | Reused barriers need a round number so late wakeups do not confuse rounds. |
| 3 | **Broadcast** | Reused: `sync.Cond.Broadcast` wakes every waiter to re-check the predicate. |

## Hint

Track `count` and `round`. The last arrival resets `count`, increments `round`, and broadcasts.

## Validate

```bash
make verify
```

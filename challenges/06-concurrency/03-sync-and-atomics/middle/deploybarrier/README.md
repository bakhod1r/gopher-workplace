# Rollout Phase Barrier

**Level:** middle
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A deploy tool rolls a change out in phases: every region must finish phase one before any region starts phase two. The coordinator makes each region's goroutine wait at a barrier that releases only when all of them have arrived, and the barrier is reused for each phase.

## Task

Implement the stubbed functions in [deploybarrier.go](deploybarrier.go) so that:

1. `NewBarrier` clamps `n` to at least 1 and attaches a `sync.Cond`.
2. `Wait` blocks until all `n` participants have arrived.
3. The last arrival resets the count, advances the phase, and releases everyone.
4. The barrier must work again for the next phase.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewBarrier(3).Phase()
Output: 0
```

**Example 2:**

```
Input:  b := NewBarrier(1); b.Wait(); b.Phase()
Output: 1
```

**Example 3:**

```
Input:  8 regions: 7 wait, the 8th releases all
Output: Phase() == 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Cyclic barrier | Count arrivals; the last one flips the state and broadcasts, then the counter resets for the next round. |
| 2 | Phase as the predicate | Waiting on "my phase is still current" is what makes the barrier reusable — an arrival counter alone would be ambiguous. |
| 3 | `Broadcast`, not `Signal` | Every participant of the phase must wake, not just one. |
| 4 | Snapshot before waiting | Read the phase **before** incrementing, so the predicate refers to the round you joined. |

## Hint

Record `phase := b.phase` on entry. The last arrival sets `arrived = 0`, `phase++`, and broadcasts; everyone else loops `for phase == b.phase { cond.Wait() }`.

## Validate

```bash
make verify
```

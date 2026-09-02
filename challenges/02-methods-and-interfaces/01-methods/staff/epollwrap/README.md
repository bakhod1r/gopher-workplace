# Epoll Interest Set

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An epoll instance is a registry: each file descriptor declares which events it
cares about, and each wakeup reports which events actually fired. The readiness
report is only useful after it has been intersected with the interest set — the
kernel may report bits nobody asked for, and for descriptors nobody registered.

## Task

Implement `Wait` on `*Epoll` in [epollwrap.go](epollwrap.go):

1. For each entry in `ready`, look up the registered interest mask.
2. Keep the descriptor only if `readyMask & interestMask != 0`.
3. Ignore descriptors that are not registered.
4. Return the surviving descriptors sorted ascending.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(3, EventRead); Wait({3: EventRead})
Output: [3]
```

**Example 2:**

```
Input:  Add(7, EventWrite); Wait({7: EventRead})
Output: []      (registered, but the wrong event fired)
```

**Example 3:**

```
Input:  Add(5, EventRead|EventWrite); Wait({5: EventWrite, 9: EventRead})
Output: [5]     (one bit overlaps; fd 9 was never registered)
```

_Explanation:_ a single overlapping bit is enough; an unregistered fd is never enough.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bitmask intersection** | `a&b != 0` tests "any shared bit", which is the readiness question. |
| 2 | **Comma-ok lookup** | An unregistered fd must be skipped, not treated as interest `0`. |
| 3 | **Deterministic output from a map** | Map iteration order is randomized, so the result must be sorted before it is comparable. |

## Hint

`for fd, mask := range ready` gives a random order every run — collect first,
`sort.Ints` last. And return an empty (or nil) slice, never a partially filled
one, when nothing matched.

## Validate

```bash
make verify
```

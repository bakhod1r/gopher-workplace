# Select Semantics

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A coordinator multiplexes several channels. Subtle `select` behaviour — nil channels, default, closed channels — decides whether it stalls or spins.

## Task

Implement the stub(s) in [chanselect.go](chanselect.go):

1. Implement `TryRecv`, a non-blocking receive returning `(value, ok, ready)`.
2. Implement `Drain`, which reads two channels until both are closed, returning the values in arrival-agnostic total order (sorted).
3. Implement `FirstReady`, which returns the first value available from either channel and reports which one.
4. Constraint: no busy-wait loops — a closed channel must be disabled by setting its variable to nil, not by spinning on it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TryRecv on an empty open channel
Output: ready false
```

**Example 2:**

```
Input:  Drain over two channels
Output: every value from both, sorted
```

**Example 3:**

```
Input:  FirstReady with only b ready
Output: b's value
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **nil channels in select** | A nil channel blocks forever, which disables that case. |
| 2 | **Closed channel semantics** | A closed channel is always ready and yields the zero value. |
| 3 | **default versus blocking** | `default` makes select non-blocking; omitting it blocks until a case fires. |

## Hint

When a receive reports `!ok`, set that channel variable to nil so its case stops firing.

## Validate

```bash
make verify
```

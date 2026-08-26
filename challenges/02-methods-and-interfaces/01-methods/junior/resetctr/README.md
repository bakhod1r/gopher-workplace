# Reset Counter

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A session tracker resets its request counter at the start of each interval.

## Task

Implement `Reset` on `*Counter` in [resetctr.go](resetctr.go):

1. Set `N` to `0`.
2. Pointer receiver — the caller must see the reset.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Counter{42}; c.Reset(); c.Value()
Output: 0
```

**Example 2:**

```
Input:  c := Counter{0}; c.Reset(); c.Value()
Output: 0
```

**Example 3:**

```
Input:  c := Counter{-5}; c.Reset(); c.Value()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `*Counter` — mutation visible to caller. |
| 2 | **Zero value** | Setting `N = 0` is the Go zero value for `int`. |

## Hint

`c.N = 0` on a `*Counter` receiver.

## Validate

```bash
make verify
```

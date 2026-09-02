# Flux Store

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

In a Flux/Redux-shaped architecture, state only changes in one place: a
dispatcher that takes an action and applies its effect. Nothing else may write
`Count` — which is what makes the state's history reconstructible from the
action log.

## Task

Implement `Dispatch` on `*Store` in [fluxpattern.go](fluxpattern.go):

1. `"INC"` increments `s.Count`.
2. `"DEC"` decrements `s.Count`.
3. Any other action leaves the state unchanged.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Dispatch("INC")
Output: Count == 1
```

**Example 2:**

```
Input:  INC, INC, DEC
Output: Count == 1
```

**Example 3:**

```
Input:  Dispatch("RESET")
Output: Count unchanged (unknown actions are ignored)
```

_Explanation:_ an unrecognized action is a no-op, not an error.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Single write path** | Every mutation goes through one method, so state transitions are auditable. |
| 2 | **`switch` on a string action** | The default arm being empty is a deliberate choice. |
| 3 | **Pointer receiver** | Without it the store's state never changes. |

## Hint

A `switch action` with two cases. Do not add a `default` that panics — unknown
actions must be ignored.

## Validate

```bash
make verify
```

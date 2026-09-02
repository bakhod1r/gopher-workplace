# Observer Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A subject owns some state and a list of interested parties. When the state
changes, every observer is told — synchronously, in the order they registered.
The subject knows nothing about what the observers do.

## Task

Implement `SetState` on `*Subject` in [observer.go](observer.go):

1. Set `s.state = val`.
2. Call every observer in `s.observers` with the new value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  one observer sum += state; SetState(10)
Output: sum == 10
```

**Example 2:**

```
Input:  observers sum += state and sum += state*2; SetState(10)
Output: sum == 30
```

**Example 3:**

```
Input:  SetState(5) on a subject with no observers
Output: state updated, nothing called
```

_Explanation:_ ranging over a nil slice is a no-op.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Function type as observer** | `type Observer func(int)` — no interface needed for a single-method contract. |
| 2 | **State first, notify second** | Observers may read the subject; they must see the new value. |
| 3 | **Pointer receiver** | The state write must be visible to the caller. |

## Hint

Assign before you notify, then a single `range` loop calling `o(val)`.

## Validate

```bash
make verify
```

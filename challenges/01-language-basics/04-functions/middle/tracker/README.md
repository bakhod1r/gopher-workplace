# Closure Sharing State

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Two closures created in the same scope share the same captured variable — a
lightweight object with methods, built from functions.

## Task

Implement `NewTracker` in [tracker.go](tracker.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  add, total := NewTracker(); add(3); add(4); total()
Output: 7
```

**Example 2:**

```
Input:  both closures share sum
Output: true
```

**Example 3:**

```
Input:  total() before any add
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared capture** | Both closures see the same `sum`. |
| 2 | **Multiple returns of funcs** | Return two function values. |
| 3 | **Encapsulation** | `sum` is private to the closures. |

## Hint

Declare `sum := 0`; return `func(n int){ sum += n }` and `func() int { return sum }`.

## Validate

```bash
make verify
```

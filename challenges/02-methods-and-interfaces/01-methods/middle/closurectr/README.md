# Closure Counter

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A sequence generator needs a "next ID" function. Instead of a global variable,
use a method value bound to a private Counter — clean, testable, no shared state.

## Task

Implement `NewCounter` in [closurectr.go](closurectr.go):

1. Create a `Counter` (starts at 0).
2. Return its `Inc` method as a bound method value.
3. Each call returns the next count: 1, 2, 3, ...
4. Multiple counters are independent.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  next := NewCounter(); next(), next(), next()
Output: 1, 2, 3
```

**Example 2:**

```
Input:  a := NewCounter(); b := NewCounter(); a(), a(), b()
Output: 1, 2, 1
```

**Example 3:**

```
Input:  next := NewCounter(); for i := 0; i < 100; i++ { next() }; next()
Output: 101
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method values** | `c.Inc` binds the pointer receiver — all calls share state. |
| 2 | **Pointer receiver** | `*Counter` ensures shared mutation across calls. |
| 3 | **Encapsulation** | The counter is hidden; only the increment function is exposed. |

## Hint

`c := &Counter{}; return c.Inc` — the method value captures the pointer.

## Validate

```bash
make verify
```

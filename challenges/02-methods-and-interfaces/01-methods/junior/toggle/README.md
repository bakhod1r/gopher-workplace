# Toggle

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A feature-flag system toggles flags at runtime. Each call flips the state.

## Task

Implement `Toggle` on `*Switch` in [toggle.go](toggle.go):

1. Flip `On` from `true` to `false` or vice versa.
2. Pointer receiver — mutation must persist.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := Switch{On: false}; s.Toggle()
Output: s.On == true
```

**Example 2:**

```
Input:  s := Switch{On: true}; s.Toggle()
Output: s.On == false
```

**Example 3:**

```
Input:  s := Switch{}; s.Toggle(); s.Toggle()
Output: s.On == false (back to start)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | In-place boolean flip. |
| 2 | **Boolean negation** | `s.On = !s.On`. |

## Hint

`s.On = !s.On` — one line.

## Validate

```bash
make verify
```

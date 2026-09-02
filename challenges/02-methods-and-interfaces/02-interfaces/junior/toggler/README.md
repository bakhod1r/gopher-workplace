# Toggler

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A settings panel flips switches that each remember their own state.

## Task

Implement the stub(s) in [toggler.go](toggler.go):

1. Implement `Toggle` and `State` on `*Switch`.
2. Implement `ToggleAll`, which toggles every element and returns how many ended up on.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := &Switch{}; s.Toggle(); s.State()
Output: true
```

**Example 2:**

```
Input:  s.Toggle(); s.State()
Output: false
```

**Example 3:**

```
Input:  ToggleAll([]Toggler{&Switch{}, &Switch{On: true}})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stateful interface** | The contract mixes a mutator and a reader. |
| 2 | **Pointer receivers** | Both methods use `*Switch` so the method set stays consistent. |
| 3 | **Boolean negation** | Reused: `s.On = !s.On`. |

## Hint

Toggling twice returns to the original state.

## Validate

```bash
make verify
```

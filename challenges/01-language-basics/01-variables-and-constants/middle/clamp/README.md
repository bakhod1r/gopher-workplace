# Clamp to Range

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Clamping needs valid bounds. When `lo > hi` you normalize them in a local
scope — a good place to practice `if`-init variables and block scope without
shadowing the parameters.

## Task

In [clamp.go](clamp.go):

1. If `lo > hi`, swap them (locally).
2. Return `v` limited to `[lo, hi]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Clamp(5, 0, 10)
Output: 5
```

**Example 2:**

```
Input:  Clamp(-3, 0, 10)
Output: 0
```

**Example 3:**

```
Input:  Clamp(5, 10, 0)
Output: 5
```

_Explanation:_ Reversed bounds are swapped first.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Block scope** | Variables declared in a block vanish at its end. |
| 2 | **Shadowing risk** | `lo, hi :=` inside a block would shadow, not update, the params. |
| 3 | **Multiple assignment** | `lo, hi = hi, lo` swaps without a temp. |

## Hint

Reassign with `lo, hi = hi, lo` (use `=`, not `:=`, or you shadow the params).

## Validate

```bash
make verify
```

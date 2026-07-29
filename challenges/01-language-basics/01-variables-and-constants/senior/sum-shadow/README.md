# Shadowed Accumulator

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

The loop looks like it accumulates into `total`, but `SumPositive` always
returns 0. A `:=` inside the `if` block creates a new `total` that dies each
iteration — the outer one never changes.

## Task

Fix the single line between the markers in [accumulate.go](accumulate.go) so the
outer `total` accumulates. Keep the signature and the surrounding code.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumPositive([1 2 3])
Output: 6
```

**Example 2:**

```
Input:  SumPositive([1 -2 3])
Output: 4
```

**Example 3:**

```
Input:  SumPositive(nil)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shadowing** | `:=` in an inner block declares a fresh variable. |
| 2 | **Block scope** | The inner `total` is discarded at the block's end. |
| 3 | **`=` vs `:=`** | Use `=` to update the existing variable. |

## Hint

`total := total + x` shadows. Change `:=` to `+=` (and drop the `_ = total`).

## Validate

```bash
make verify
```

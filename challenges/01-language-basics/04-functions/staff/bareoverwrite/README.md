# Named Return Left Unset

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A bare `return` sends the named result's CURRENT value. Since `result` was
never assigned, it is 0; the deferred `result *= 2` then doubles 0. Assign
`result` (from `local`) before returning so the defer doubles x.

## Task

Fix [bareoverwrite.go](bareoverwrite.go) so `Doubled(21)` returns 42.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Doubled(21)
Output: 42
```

**Example 2:**

```
Input:  Doubled(0)
Output: 0
```

**Example 3:**

```
Input:  Doubled(-5)
Output: -10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bare return uses current named values** | Unset named results are their zero value. |
| 2 | **Deferred adjust** | `result *= 2` runs after the return assigns. |
| 3 | **Assign before returning** | Set `result` so the defer has something to double. |

## Hint

Assign the named return before the bare return: `result = local; return`.

## Validate

```bash
make verify
```

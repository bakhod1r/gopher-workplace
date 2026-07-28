# Max of Three

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _functions-basics_

## Context

Combining comparisons is the bread-and-butter of small helper functions.

## Task

Implement `Max3` in [max3.go](max3.go).

Do **not** change the function signature or the tests.

## Examples

```go
Max3(1, 2, 3)    // => 3
Max3(5, 9, 5)    // => 9
Max3(-1, -2, -3) // => -1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Comparison chaining** | Track a running maximum. |
| 2 | **Negatives** | Seed from an argument, never from 0. |
| 3 | **Single result** | One `int`. |

## Hint

Start `m := a`, then `if b > m { m = b }` and the same for `c`.

## Validate

```bash
make verify
```

# Unsigned Underflow

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`have - sold` on `uint` wraps around when `sold > have`, producing a gigantic
number instead of 0. The zero value and unsigned wraparound combine into a
classic inventory bug.

## Task

Fix the single line between the markers in [stock.go](stock.go) so overselling
returns 0, never a wrapped value.

## Examples

```go
Remaining(10, 3) // => 7
Remaining(5, 5)  // => 0
Remaining(2, 9)  // => 0  (not 18446744073709551609)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Unsigned wraparound** | `uint` has no negatives; `2-9` wraps to a huge value. |
| 2 | **Guard before subtract** | Compare `sold >= have` first. |
| 3 | **Zero value** | The clamped result is the `uint` zero. |

## Hint

Guard it: `if sold >= have { return 0 }` then `return have - sold`.

## Validate

```bash
make verify
```

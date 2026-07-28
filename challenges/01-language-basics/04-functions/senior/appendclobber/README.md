# Append Clobbers Shared Cap

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Because `base` has spare capacity, `append(base, x)` writes into base's backing
array and returns a slice sharing it. The second `append(base, y)` writes to the
same slot, clobbering the first result. Clip the base so each append reallocates.

## Task

Fix [appendclobber.go](appendclobber.go) so `a` and `b` are independent.

Do **not** change the function signature or the tests.

## Examples

```go
TwoTails([1 2]cap8, 100, 200) // a=[1 2 100], b=[1 2 200] independent
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Append into spare capacity** | It reuses and shares the array. |
| 2 | **Three-index slice** | `base[:len:len]` caps capacity, forcing a copy. |
| 3 | **Independent results** | Each append must own its array. |

## Hint

Clip the base to its length before appending: `b2 := base[:len(base):len(base)]`, then append `x` and `y` to `b2` separately (each reallocates).

## Validate

```bash
make verify
```

# Two Appends Share Capacity

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

`append(base, x)` and `append(base, y)` both write into `base`'s spare capacity at
the same index, so the second overwrites the first — `b` ends up `[1 2 4]`.

## Task

Fix the base between the markers in
[appendtwobranch.go](appendtwobranch.go) so the two appends don't alias.

## Examples

```go
b, c := Branch([1,2](cap>2), 3, 4) // b=[1 2 3], c=[1 2 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared spare cap** | Both appends target the same slot. |
| 2 | **Clip capacity** | `a[:len(a):len(a)]` forces realloc. |
| 3 | **Independent results** | Each append gets its own array. |

## Hint

`base := a[:len(a):len(a)]`.

## Validate

```bash
make verify
```

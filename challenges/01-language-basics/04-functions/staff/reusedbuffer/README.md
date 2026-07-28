# Returning a Reused Buffer

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

The closure reuses one backing array (`buf[:0]` keeps the capacity). Returning
`buf` hands out a view into that shared array, so the NEXT call's `buf[:0]` +
append overwrites the elements an earlier caller still holds. Return a copy.

## Task

Fix [reusedbuffer.go](reusedbuffer.go) so earlier results survive later calls.

Do **not** change the function signature or the tests.

## Examples

```go
r := Reader(); a := r(1,2); b := r(3,4) // a stays [1 2]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Reused backing array** | `buf[:0]` keeps capacity and reuses memory. |
| 2 | **Returning a shared view** | The result aliases the buffer. |
| 3 | **Defensive copy** | Copy out so callers own their data. |

## Hint

Return an independent copy: `return append([]int(nil), buf...)` (or make+copy).

## Validate

```bash
make verify
```

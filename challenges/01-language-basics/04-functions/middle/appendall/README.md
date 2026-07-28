# Forward Variadic to Append

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

Spreading a variadic slice into `append` (`append(dst, extra...)`) forwards all
elements. To avoid mutating `base`, copy it first.

## Task

Implement `Concat` in [appendall.go](appendall.go) without mutating `base`.

Do **not** change the function signature or the tests.

## Examples

```go
Concat([]int{1,2}, 3,4,5) // => [1 2 3 4 5]
Concat(nil, 1)            // => [1]
Concat([]int{9})          // => [9]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic spread into append** | `append(dst, extra...)`. |
| 2 | **Copy to avoid aliasing** | `append(base, ...)` may reuse base's array. |
| 3 | **Fresh result** | Callers keep their original. |

## Hint

Make a copy: `out := append([]int(nil), base...)`, then `return append(out, extra...)`.

## Validate

```bash
make verify
```

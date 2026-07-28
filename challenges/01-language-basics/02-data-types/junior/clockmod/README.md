# 24-Hour Clock

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Adding hours wraps around midnight. Because Go's `%` can be negative, wrapping
backward needs an extra step to stay in `0..23`.

## Task

Implement `AddHours(h, add)` returning the resulting hour in `0..23`. `add` may
be negative or larger than 24.

## Examples

```go
AddHours(23, 1)  // => 0
AddHours(0, -1)  // => 23
AddHours(6, -30) // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modulo wrap** | `(h+add) % 24` folds into range. |
| 2 | **Negative remainder** | Go's `%` can be negative; normalize it. |
| 3 | **Normalize** | `((x % 24) + 24) % 24` is always 0..23. |

## Hint

`((h+add)%24 + 24) % 24` guarantees a non-negative result.

## Validate

```bash
make verify
```

# Join with Separator

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The inverse of split: glue parts with a separator, but not before the first.

## Task

Implement `Join(parts, sep)` without `strings.Join`.

## Examples

```go
Join([]string{"a","b","c"}, ",") // => "a,b,c"
Join([]string{"x"}, ",")         // => "x"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **strings.Builder** | Efficient string assembly. |
| 2 | **Separator placement** | Before all but the first. |
| 3 | **Empty case** | No parts → "". |

## Hint

For `i, p := range parts`: if `i > 0` write `sep`, then write `p`.

## Validate

```bash
make verify
```

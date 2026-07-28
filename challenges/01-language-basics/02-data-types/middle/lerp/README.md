# Linear Interpolation

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Lerp blends two values by a factor `t`: at `t=0` you get `a`, at `t=1` you get
`b`.

## Task

Implement `Lerp(a, b, t)` = `a + (b-a)*t`.

## Examples

```go
Lerp(0, 10, 0.5)  // => 5
Lerp(2, 4, 0.25)  // => 2.5
Lerp(0, 10, 1)    // => 10
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Interpolation formula** | `a + (b-a)*t`. |
| 2 | **Float arithmetic** | All operands float64. |
| 3 | **Endpoints** | t=0→a, t=1→b exactly. |

## Hint

`return a + (b-a)*t`.

## Validate

```bash
make verify
```

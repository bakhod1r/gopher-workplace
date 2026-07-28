# Order Total

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A cart totals line items: each `Order` has a price and quantity; sum
`Price*Qty`.

## Task

Implement `Total(orders)`.

## Examples

```go
Total([]Order{{"pen",150,2},{"pad",300,1}}) // => 600
Total(nil)                                   // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct fields** | Access `o.Price`, `o.Qty`. |
| 2 | **Slice of structs** | Range yields a copy of each struct. |
| 3 | **Accumulate** | Sum `Price*Qty`. |

## Hint

`for _, o := range orders { total += o.Price * o.Qty }`.

## Validate

```bash
make verify
```

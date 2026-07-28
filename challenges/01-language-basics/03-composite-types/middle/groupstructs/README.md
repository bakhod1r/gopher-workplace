# Total by Customer

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Aggregating records into per-key totals — a group-by-sum.

## Task

Implement `TotalByCustomer(orders)`.

## Examples

```go
TotalByCustomer([{ann,100},{bob,50},{ann,25}]) // => {ann:125, bob:50}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Group-by-sum** | `m[key] += value`. |
| 2 | **Map zero value** | Missing key starts at 0. |
| 3 | **Struct fields** | Read Customer, Amount. |

## Hint

`for _, o := range orders { m[o.Customer] += o.Amount }`.

## Validate

```bash
make verify
```

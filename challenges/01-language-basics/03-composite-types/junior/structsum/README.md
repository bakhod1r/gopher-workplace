# Order Total

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A cart totals line items: each `Order` has a price and quantity; sum
`Price*Qty`.

## Task

Implement `Total(orders)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [{pen,150,2},{pad,300,1},{ink,500,3}]
Output: 2100
```

_Explanation:_ 300+300+1500.

**Example 2:**

```
Input:  nil
Output: 0
```

**Example 3:**

```
Input:  [{a,100,1}]
Output: 100
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

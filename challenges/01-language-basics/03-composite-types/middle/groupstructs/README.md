# Total by Customer

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Aggregating records into per-key totals — a group-by-sum.

## Task

Implement `TotalByCustomer(orders)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [{ann,100},{bob,50},{ann,25},{bob,75}]
Output: {ann:125, bob:125}
```

**Example 2:**

```
Input:  [{ann,10}]
Output: {ann:10}
```

**Example 3:**

```
Input:  []
Output: {}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

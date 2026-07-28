# Mutating the Range Copy

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`for _, o := range orders` gives a **copy** of each struct. `o.Price -= ...`
changes the copy; the slice is untouched.

## Task

Fix the loop between the markers in
[rangecopymutate.go](rangecopymutate.go) to mutate the slice element.

## Examples

```go
Discount([]Order{{100}}, 10) // orders[0].Price becomes 90
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Range copies** | The value variable is a copy. |
| 2 | **Index to mutate** | `orders[i].Price = ...`. |
| 3 | **Structs are values** | Copied on assignment/range. |

## Hint

`for i := range orders { orders[i].Price -= orders[i].Price * pct / 100 }`.

## Validate

```bash
make verify
```

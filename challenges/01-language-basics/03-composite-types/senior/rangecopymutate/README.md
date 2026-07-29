# Mutating the Range Copy

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`for _, o := range orders` gives a **copy** of each struct. `o.Price -= ...`
changes the copy; the slice is untouched.

## Task

Fix the loop between the markers in
[rangecopymutate.go](rangecopymutate.go) to mutate the slice element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  orders=[{100},{200}], pct=10
Output: [{90},{180}]
```

**Example 2:**

```
Input:  orders=[{50}], pct=0
Output: [{50}]
```

**Example 3:**

```
Input:  orders=[{100}], pct=100
Output: [{0}]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
